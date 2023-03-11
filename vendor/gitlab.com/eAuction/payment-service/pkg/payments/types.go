package payments

const (
	TxStatusPending   = "PENDING"
	TxStatusOk        = "OK"
	TxStatusConfirmed = "CONFIRMED"
	TxStatusReversed  = "REVERSED"
	TxStatusRefunded  = "REFUNDED"
	TxStatusTimeOut   = "TIMEOUT"
	TxStatusFailed    = "FAILED"
	TxStatusDeclined  = "DECLINED"
)

const (
	TxPurposeDeposit    = "DEPOSIT"
	TxPurposeServiceFee = "SERVICE_FEE"
)
