package types

//go:generate jsonenums -tprefix=false -transform=snake -type=DepositOptionState
type DepositOptionState int32

const (
	//DepositOptionStateTaken - deposit option is already selected for this lot
	DepositOptionStateTaken DepositOptionState = iota + 1
	//DepositOptionStateNotAvailable - due to reasons it's not allowed to select specified option
	DepositOptionStateNotAvailable
	//DepositOptionStateAvailable - deposit option can be selected
	DepositOptionStateAvailable
)
