package network

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	hbls "github.com/herumi/bls-eth-go-binary/bls"
	"github.com/pkg/errors"
	"github.com/protolambda/zrnt/eth2/beacon/common"
	"github.com/protolambda/zrnt/eth2/util/hashing"
	"github.com/protolambda/ztyp/tree"
	util "github.com/wealdtech/go-eth2-util"
	"strings"
)

func GenerateDepositData(forkVersion string, accountMax uint64, accountMin uint64, amountGwei uint64, validatorsMnemonic string, withdrawalsMnemonic string, asJsonList bool) (string, error) {
	var genesisForkVersion common.Version
	err := genesisForkVersion.UnmarshalText([]byte(forkVersion))
	if err != nil {
		return "", errors.Wrap(err, "cannot decode fork version")
	}

	valSeed, err := mnemonicToSeed(validatorsMnemonic)
	if err != nil {
		return "", errors.Wrap(err, "bad validator mnemonic")
	}
	withdrSeed, err := mnemonicToSeed(withdrawalsMnemonic)
	if err != nil {
		return "", errors.Wrap(err, "bad validator mnemonic")
	}

	var output strings.Builder

	if asJsonList {
		output.WriteString("[")
	}
	for i := accountMin; i < accountMax; i++ {
		valAccPath := fmt.Sprintf("m/12381/3600/%d/0/0", i)
		val, err := util.PrivateKeyFromSeedAndPath(valSeed, valAccPath)
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("failed to create validator private key for path %q", valAccPath))
		}
		withdrAccPath := fmt.Sprintf("m/12381/3600/%d/0", i)
		withdr, err := util.PrivateKeyFromSeedAndPath(withdrSeed, withdrAccPath)
		if err != nil {
			return "", errors.Wrap(err, fmt.Sprintf("failed to create validator private key for path %q", valAccPath))
		}

		var pub common.BLSPubkey
		copy(pub[:], val.PublicKey().Marshal())

		var withdrPub common.BLSPubkey
		copy(withdrPub[:], withdr.PublicKey().Marshal())
		withdrCreds := hashing.Hash(withdrPub[:])
		withdrCreds[0] = common.BLS_WITHDRAWAL_PREFIX

		data := common.DepositData{
			Pubkey:                pub,
			WithdrawalCredentials: withdrCreds,
			Amount:                common.Gwei(amountGwei),
			Signature:             common.BLSSignature{},
		}
		msgRoot := data.ToMessage().HashTreeRoot(tree.GetHashFn())
		var secKey hbls.SecretKey
		err = secKey.Deserialize(val.Marshal())
		if err != nil {
			return "", errors.Wrap(err, "cannot convert validator priv key")
		}

		dom := common.ComputeDomain(common.DOMAIN_DEPOSIT, genesisForkVersion, common.Root{})
		msg := common.ComputeSigningRoot(msgRoot, dom)
		sig := secKey.SignHash(msg[:])
		copy(data.Signature[:], sig.Serialize())

		dataRoot := data.HashTreeRoot(tree.GetHashFn())
		jsonData := map[string]interface{}{
			"account":                valAccPath, // for ease with tracking where it came from.
			"pubkey":                 hex.EncodeToString(data.Pubkey[:]),
			"withdrawal_credentials": hex.EncodeToString(data.WithdrawalCredentials[:]),
			"signature":              hex.EncodeToString(data.Signature[:]),
			"value":                  uint64(data.Amount),
			"deposit_data_root":      hex.EncodeToString(dataRoot[:]),
			"version":                1, // ethereal cli requirement
		}
		jsonStr, err := json.Marshal(jsonData)
		if asJsonList && i+1 < accountMax {
			jsonStr = append(jsonStr, ',')
		}
		if err != nil {
			return "", errors.Wrap(err, "could not encode deposit data to json")
		}
		output.WriteString(string(jsonStr))
	}
	if asJsonList {
		output.WriteString("]")
	}

	return output.String(), nil
}
