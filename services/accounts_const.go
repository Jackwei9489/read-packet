package services

// TransferStatus 转账状态
type TransferStatus int8

const (
	// TransferredStatusFailed 转账失败
	TransferredStatusFailed TransferStatus = -1
	// TransferredStatusInSufficient 余额不足
	TransferredStatusInSufficient TransferStatus = 0
	// TransferredStatusSuccess 转账成功
	TransferredStatusSuccess TransferStatus = 1
)

// ChangeType 转账的类型
type ChangeType int8

const (
	// AccountCreated 账户创建
	AccountCreated ChangeType = 0
	// AccountStoreValue 账户储值
	AccountStoreValue ChangeType = 1
	// EnvelopeExpense 资金支出
	EnvelopeExpense ChangeType = -2
	// EnvelopeIncoming 资金收入
	EnvelopeIncoming ChangeType = 2
	// EnvelopeExpiredRefund 过期退款
	EnvelopeExpiredRefund ChangeType = 3
)

// ChangeFlag 资金交易变化标识
type ChangeFlag int8

const (
	// FlagAccountCreated 创建账户=0
	FlagAccountCreated ChangeFlag = 0
	// FlagTransferOut 支出=-1
	FlagTransferOut ChangeFlag = -1
	// FlagTransferIn 收入=1
	FlagTransferIn ChangeFlag = 1
)
