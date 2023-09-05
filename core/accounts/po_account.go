package accounts

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"time"
)

type Base struct {
	Version   int       `gorm:"column:version"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

// Account 账户持久化对象
type Account struct {
	Id           int64           `gorm:"column:id;primaryKey;autoIncrement"`
	AccountNo    string          `gorm:"column:account_no;uniqueIndex;size:30"`
	AccountName  string          `gorm:"column:account_name;size:50"`
	AccountType  int             `gorm:"column:account_type"`
	CurrencyCode string          `gorm:"column:currency_code;size:10"`
	UserId       string          `gorm:"column:user_id"`
	Username     sql.NullString  `gorm:"column:username;size:50"`
	Balance      decimal.Decimal `gorm:"column:balance;type:decimal(30,6)"`
	Status       int             `gorm:"column:status"`
	Base         `gorm:"embedded"`
}
