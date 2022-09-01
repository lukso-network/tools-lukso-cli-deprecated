package network

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/protolambda/go-keystorev4"
	e2types "github.com/wealdtech/go-eth2-types/v2"
	util "github.com/wealdtech/go-eth2-util"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type PrysmAccountStore struct {
	PrivateKeys [][]byte `json:"private_keys"`
	PublicKeys  [][]byte `json:"public_keys"`
}

// Following EIP 2335
type KeyFile struct {
	id        uuid.UUID
	name      string
	publicKey e2types.PublicKey
	secretKey e2types.PrivateKey
}

type KeyEntry struct {
	KeyFile
	passphrase string
	insecure   bool
}

type WalletOutput interface {
	InsertAccount(priv e2types.PrivateKey, insecure bool, idx uint64) error
}

type WalletWriter struct {
	sync.RWMutex
	entries []*KeyEntry
}

func NewWalletWriter(entries uint64) *WalletWriter {
	return &WalletWriter{
		entries: make([]*KeyEntry, entries),
	}
}

func GenerateKeystores(accountMax uint64, accountMin uint64, sourceMnemonic string,
	insecure bool, outputDataPath string, prysmPass string) error {
	ww := NewWalletWriter(accountMax - accountMin)
	err := selectVals(sourceMnemonic, accountMin, accountMax, ww, insecure)
	if err != nil {
		return errors.Wrap(err, "failed to assign validators")
	}
	err = ww.WriteOutputs(outputDataPath, prysmPass)
	if err != nil {
		return errors.Wrap(err, "failed to write output")
	}
	return nil
}

func (ww *WalletWriter) InsertAccount(priv e2types.PrivateKey, insecure bool, idx uint64) error {
	key, err := NewKeyEntry(priv, insecure)
	if err != nil {
		return err
	}
	ww.RWMutex.Lock()
	defer ww.RWMutex.Unlock()
	ww.entries[idx] = key
	return nil
}

func NewKeyEntry(priv e2types.PrivateKey, insecure bool) (*KeyEntry, error) {
	var pass [32]byte
	n, err := rand.Read(pass[:])
	if err != nil {
		return nil, err
	}
	if n != 32 {
		return nil, errors.New("could not read sufficient secure random bytes")
	}
	// Convert it to human readable characters, to keep it manageable
	passphrase := base64.URLEncoding.EncodeToString(pass[:])
	return &KeyEntry{
		KeyFile: KeyFile{
			id:        uuid.New(),
			name:      "val_" + hex.EncodeToString(priv.PublicKey().Marshal()),
			publicKey: priv.PublicKey(),
			secretKey: priv,
		},
		passphrase: passphrase,
		insecure:   insecure,
	}, nil
}

// Narrow pubkeys: we don't want 0xAb... to be different from ab...
func narrowedPubkey(pub string) string {
	return strings.TrimPrefix(strings.ToLower(pub), "0x")
}

func selectVals(sourceMnemonic string,
	minAcc uint64, maxAcc uint64,
	output WalletOutput, insecure bool) error {

	valSeed, err := mnemonicToSeed(sourceMnemonic)
	if err != nil {
		return err
	}

	var g errgroup.Group
	// Try look for unassigned accounts in the wallet
	for i := minAcc; i < maxAcc; i++ {
		idx := i
		g.Go(func() error {
			valAccPath := fmt.Sprintf("m/12381/3600/%d/0/0", idx)
			a, err := util.PrivateKeyFromSeedAndPath(valSeed, valAccPath)
			if err != nil {
				return fmt.Errorf("account %s cannot be derived, continuing to next account", valAccPath)
			}
			pubkey := narrowedPubkey(hex.EncodeToString(a.PublicKey().Marshal()))
			if err := output.InsertAccount(a, insecure, idx-minAcc); err != nil {
				if err.Error() == fmt.Sprintf("account with name \"%s\" already exists", pubkey) {
					fmt.Printf("Account with pubkey %s already exists in output wallet, skipping it\n", pubkey)
				} else {
					return fmt.Errorf("failed to import account with pubkey %s into output wallet: %v", pubkey, err)
				}
			}

			return nil
		})

	}
	return g.Wait()
}

func (ww *WalletWriter) WriteOutputs(fpath string, prysmPass string) error {
	if _, err := os.Stat(fpath); !os.IsNotExist(err) {
		return errors.New("output for assignments already exists! Aborting")
	}
	if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
		return err
	}
	// What lighthouse requires as file name
	lighthouseKeyfileName := "voting-keystore.json"
	lighthouseKeyfilesPath := filepath.Join(fpath, "keys")
	if err := os.Mkdir(lighthouseKeyfilesPath, os.ModePerm); err != nil {
		return err
	}
	// nimbus has different keystore names
	nimbusKeyfileName := "keystore.json"
	nimbusKeyfilesPath := filepath.Join(fpath, "nimbus-keys")
	if err := os.Mkdir(nimbusKeyfilesPath, os.ModePerm); err != nil {
		return err
	}
	// teku does not nest their keystores
	tekuKeyfilesPath := filepath.Join(fpath, "teku-keys")
	if err := os.Mkdir(tekuKeyfilesPath, os.ModePerm); err != nil {
		return err
	}

	var g errgroup.Group
	// For all: write JSON keystore files, each in their own directory (lighthouse requirement)
	for _, k := range ww.entries {
		e := k
		g.Go(func() error {
			dat, err := e.MarshalJSON()
			if err != nil {
				return err
			}
			{
				// lighthouse
				keyDirPath := filepath.Join(lighthouseKeyfilesPath, e.PubHex())
				if err := os.MkdirAll(keyDirPath, os.ModePerm); err != nil {
					return err
				}
				if err := ioutil.WriteFile(filepath.Join(keyDirPath, lighthouseKeyfileName), dat, 0644); err != nil {
					return err
				}
			}
			{
				// nimbus
				keyDirPath := filepath.Join(nimbusKeyfilesPath, e.PubHex())
				if err := os.MkdirAll(keyDirPath, os.ModePerm); err != nil {
					return err
				}
				if err := ioutil.WriteFile(filepath.Join(keyDirPath, nimbusKeyfileName), dat, 0644); err != nil {
					return err
				}
			}
			{
				// teku
				if err := ioutil.WriteFile(filepath.Join(tekuKeyfilesPath, e.PubHex()+".json"), dat, 0644); err != nil {
					return err
				}
			}
			return nil
		})

	}
	{
		// For Lighthouse: they need a directory that maps pubkey to passwords, one per file
		secretsDirPath := filepath.Join(fpath, "secrets")
		if err := os.Mkdir(secretsDirPath, os.ModePerm); err != nil {
			return err
		}
		for _, k := range ww.entries {
			e := k
			g.Go(func() error {
				pubHex := e.PubHex()
				return ioutil.WriteFile(path.Join(secretsDirPath, pubHex), []byte(e.passphrase), 0644)
			})
		}
	}

	{
		// For Teku: they need a directory that maps name of keystore dir to name of secret file, but secret files end with `.txt`
		secretsDirPath := filepath.Join(fpath, "teku-secrets")
		if err := os.Mkdir(secretsDirPath, os.ModePerm); err != nil {
			return err
		}
		for _, k := range ww.entries {
			e := k
			g.Go(func() error {
				pubHex := e.PubHex()
				return ioutil.WriteFile(filepath.Join(secretsDirPath, pubHex+".txt"), []byte(e.passphrase), 0644)
			})

		}
	}

	{
		// For Lodestar: they need a directory that maps pubkey to passwords, one per file, but no 0x prefix.
		secretsDirPath := filepath.Join(fpath, "lodestar-secrets")
		if err := os.Mkdir(secretsDirPath, os.ModePerm); err != nil {
			return err
		}
		for _, k := range ww.entries {
			e := k
			g.Go(func() error {
				pubHex := e.PubHexBare()
				return ioutil.WriteFile(filepath.Join(secretsDirPath, "0x"+pubHex), []byte(e.passphrase), 0644)
			})

		}
	}

	// In general: a list of pubkeys.
	pubkeys := make([]string, 0)
	for _, e := range ww.entries {
		pubkeys = append(pubkeys, e.PubHex())
	}
	pubsData, err := json.Marshal(pubkeys)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(fpath, "pubkeys.json"), pubsData, 0644); err != nil {
		return err
	}

	// For Prysm: write outputs as a wallet and a configuration
	if err := ww.buildPrysmWallet(filepath.Join(fpath, "prysm"), prysmPass); err != nil {
		return err
	}
	return g.Wait()
}

func (ww *WalletWriter) buildPrysmWallet(outPath string, prysmPass string) error {
	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		return err
	}
	// Prysm wallet expects the following structure, assuming
	// the output path is called `prysm`:
	//  direct/
	//    accounts/
	//      all-accounts.keystore.json
	//      - Prysm doesn't know what individual keystores are, only allowing you to import them with CLI, but not simply load them as accounts.
	//      - All pubkeys/privkeys are put in two lists, encoded as JSON, and those bytes are then encrypted exactly like a single private key would be normally
	//      - And then persisted in "all-accounts.keystore.json"
	//  keymanageropts.json
	//  - '{"direct_eip_version": "EIP-2335"}'
	accountsKeystorePath := filepath.Join(outPath, "direct", "accounts")
	if err := os.MkdirAll(accountsKeystorePath, os.ModePerm); err != nil {
		return err
	}
	store := PrysmAccountStore{}
	for _, e := range ww.entries {
		store.PublicKeys = append(store.PublicKeys, e.publicKey.Marshal())
		store.PrivateKeys = append(store.PrivateKeys, e.secretKey.Marshal())
	}
	storeBytes, err := json.MarshalIndent(&store, "", "\t")
	if err != nil {
		return err
	}

	kdfParams, err := keystorev4.NewPBKDF2Params()
	if err != nil {
		return fmt.Errorf("failed to create PBKDF2 params: %w", err)
	}
	cipherParams, err := keystorev4.NewAES128CTRParams()
	if err != nil {
		return fmt.Errorf("failed to create AES128CTR params: %w", err)
	}
	crypto, err := keystorev4.Encrypt(storeBytes, []byte(prysmPass),
		kdfParams, keystorev4.Sha256ChecksumParams, cipherParams)
	if err != nil {
		return fmt.Errorf("failed to encrypt secret: %w", err)
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	keystore := &keystorev4.Keystore{
		Crypto:      *crypto,
		Description: "",
		Pubkey:      nil,
		Path:        "",
		UUID:        id,
		Version:     4,
	}
	encodedStore, err := json.MarshalIndent(keystore, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(accountsKeystorePath, "all-accounts.keystore.json"), encodedStore, 0644); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(outPath, "keymanageropts.json"), []byte(`{"direct_eip_version": "EIP-2335"}`), 0644); err != nil {
		return err
	}
	return nil
}

func (ke *KeyEntry) MarshalJSON() ([]byte, error) {
	var salt [32]byte
	if _, err := rand.Read(salt[:]); err != nil {
		return nil, err
	}
	kdfParams := &keystorev4.PBKDF2Params{
		Dklen: 32,
		C:     262144,
		Prf:   "hmac-sha256",
		Salt:  salt[:],
	}
	if ke.insecure { // INSECURE but much faster, this is useful for ephemeral testnets
		kdfParams.C = 2
	}
	cipherParams, err := keystorev4.NewAES128CTRParams()
	if err != nil {
		return nil, fmt.Errorf("failed to create AES128CTR params: %w", err)
	}
	crypto, err := keystorev4.Encrypt(ke.secretKey.Marshal(), []byte(ke.passphrase),
		kdfParams, keystorev4.Sha256ChecksumParams, cipherParams)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt secret: %w", err)
	}
	keystore := &keystorev4.Keystore{
		Crypto:      *crypto,
		Description: fmt.Sprintf("0x%x", ke.publicKey.Marshal()),
		Pubkey:      ke.publicKey.Marshal(),
		Path:        "",
		UUID:        ke.id,
		Version:     4,
	}
	return json.Marshal(keystore)
}

func (ke *KeyEntry) PubHex() string {
	return "0x" + hex.EncodeToString(ke.publicKey.Marshal())
}

func (ke *KeyEntry) PubHexBare() string {
	return hex.EncodeToString(ke.publicKey.Marshal())
}
