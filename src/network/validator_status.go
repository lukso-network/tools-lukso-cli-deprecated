package network

type ValidatorStatus int64

const (
	ValidatorStatusNoDeposit ValidatorStatus = iota
	ValidatorStatusUnknown
	ValidatorStatusPending
	ValidatorStatusActive
	ValidatorStatusExited
)

func (v ValidatorStatus) String() string {
	switch v {
	case ValidatorStatusNoDeposit:
		return "no deposit"
	case ValidatorStatusUnknown:
		return "unknown"
	case ValidatorStatusPending:
		return "pending"
	case ValidatorStatusActive:
		return "active"
	case ValidatorStatusExited:
		return "exited"
	default:
		return "unknown"
	}
}
