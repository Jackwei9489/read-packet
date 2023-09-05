package accounts

import (
	"github.com/shopspring/decimal"
	"red-packet/services"
)

type AccountLog struct {
	Id              int64               `gorm:"column:id,primaryKey,autoIncrement"`
	LogNo           string              `gorm:"column:log_no,uniqueIndex"`
	AccountNo       string              `gorm:"column:account_no"`
	UserId          string              `gorm:"column:user_id"`
	Username        string              `gorm:"column:username"`
	TargetAccountNo string              `gorm:"column:target_account_no"`
	TargetUserId    string              `gorm:"column:target_user_id"`
	TargetUserName  string              `gorm:"column:target_username"`
	Amount          decimal.Decimal     `gorm:"column:amount"`
	Balance         decimal.Decimal     `gorm:"column:balance"`
	ChangeType      services.ChangeType `gorm:"column:change_type"`
	ChangeFlag      services.ChangeFlag `gorm:"column:change_flag"`
	Status          int                 `gorm:"column:status"`
	Desc            string              `gorm:"column:desc"`
	base            Base                `gorm:"embedded"`
}
