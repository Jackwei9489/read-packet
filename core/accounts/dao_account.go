package accounts

import (
	"fmt"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountDao struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

func (dao *AccountDao) getFirst(query *Account) *Account {
	a := &Account{}
	result := dao.db.Where(query).First(a)
	if result.Error != nil {
		// 查询出错
		dao.logger.Error(result.Error)
		return nil
	}
	return a
}

// GetByAccountNo 根据AccountNo 返回一条数据
func (dao *AccountDao) GetByAccountNo(accountNo string) *Account {
	a := &Account{
		AccountNo: accountNo,
	}
	return dao.getFirst(a)
}

// GetByUserId 通过用户Id和账户类型来查询账户信息
func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {
	a := &Account{
		UserId:      userId,
		AccountType: accountType,
	}
	return dao.getFirst(a)
}

// Insert 账户数据的插入
func (dao *AccountDao) Insert(a *Account) (int64, error) {
	result := dao.db.Save(a)
	if result.Error != nil {
		return 0, result.Error
	}
	return a.Id, nil
}

// UpdateBalance 账户余额的更新
func (dao *AccountDao) UpdateBalance(accountNo string,
	amount decimal.Decimal) (rows int64, err error) {
	account := dao.GetByAccountNo(accountNo)
	if account == nil {
		return 0, fmt.Errorf("更新账户余额->账户不存在")
	}
	result := dao.db.Model(&Account{}).
		Where("account_no=? and version=? and balance + CAST(? AS DECIMAL(30,6)) >= 0", accountNo, account.Version, amount).
		Updates(map[string]any{
			"version": gorm.Expr("version + ?", 1),
			"balance": gorm.Expr("balance + CAST(? AS DECIMAL(30,6))", amount),
		})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// UpdateStatus 账户状态的更新
func (dao *AccountDao) UpdateStatus(accountNo string, status int) (rows int64, err error) {
	account := dao.GetByAccountNo(accountNo)
	if account == nil {
		return 0, fmt.Errorf("更新账户状态->账户不存在")
	}
	result := dao.db.Model(&Account{}).
		Where("accountNo=? and version=?", accountNo, account.Version).
		Updates(map[string]any{
			"version": gorm.Expr("version + ?", 1),
			"status":  status,
		})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}
