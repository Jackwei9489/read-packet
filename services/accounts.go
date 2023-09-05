package services

import "time"

type AccountService interface {
	CreateAccount(command AccountCreatedDTO) (*AccountDTO, error)
	Transfer(command AccountTransferDTO) (TransferStatus, error)
	StoreValue(command AccountTransferDTO) (TransferStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
}

// TradeParticipator 账户交易参与者
type TradeParticipator struct {
	AccountNo string
	UserId    string
	UserName  string
}

// AccountTransferDTO 账户转账
type AccountTransferDTO struct {
	TradeNo     string
	TradeBody   TradeParticipator
	TradeTarget TradeParticipator
	AmountStr   string
	ChangeType  ChangeType
	ChangeFlag  ChangeFlag
	Desc        string
}

// AccountCreatedDTO 账户创建
type AccountCreatedDTO struct {
	UserId       string
	UserName     string
	AccountName  string
	AccountType  int
	CurrencyCode string
	Amount       string
}

type AccountDTO struct {
	AccountCreatedDTO
	AccountNo string
	CreatedAt time.Time
}
