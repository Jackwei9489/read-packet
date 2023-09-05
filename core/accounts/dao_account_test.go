package accounts

import (
	"database/sql"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"red-packet/infra/base"
	_ "red-packet/testx"
	"testing"
)

var (
	kd     = ksuid.New()
	db     = base.DataBase()
	logger = base.Logger()
)

func TestAccountDao_GetByAccountNo(t *testing.T) {
	tx := db.Begin()
	dao := &AccountDao{
		db:     tx,
		logger: logger,
	}
	Convey("通过账户编号查询账户数据", t, func() {
		So(tx.Error, ShouldBeNil)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
			tx.Rollback()
		}()
		a := &Account{
			AccountNo:   kd.Next().String(),
			AccountName: "测试资金账户",
			UserId:      kd.Next().String(),
			Username:    sql.NullString{String: "测试用户", Valid: true},
			Balance:     decimal.NewFromFloat(100),
			Status:      1,
		}
		id, err := dao.Insert(a)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		account := dao.GetByAccountNo(a.AccountNo)
		So(account, ShouldNotBeNil)
		So(account.Balance.String(), ShouldEqual, a.Balance.String())
		So(account.CreatedAt, ShouldNotBeNil)
		So(account.UpdatedAt, ShouldNotBeNil)
	})
}

func TestAccountDao_GetByUserId(t *testing.T) {
	tx := db.Begin()
	dao := &AccountDao{
		db:     tx,
		logger: logger,
	}
	Convey("通过用户ID和账户类型查询账户数据", t, func() {
		So(tx.Error, ShouldBeNil)
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
			tx.Rollback()
		}()
		a := &Account{
			AccountNo:   kd.Next().String(),
			AccountName: "测试资金账户",
			UserId:      kd.Next().String(),
			Username:    sql.NullString{String: "测试用户", Valid: true},
			Balance:     decimal.NewFromFloat(100),
			Status:      1,
			AccountType: 2,
		}
		id, err := dao.Insert(a)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		account := dao.GetByUserId(a.UserId, a.AccountType)
		So(account, ShouldNotBeNil)
		So(account.UserId, ShouldEqual, a.UserId)
		So(account.AccountType, ShouldEqual, a.AccountType)
	})

}

func TestAccountDao_UpdateBalance(t *testing.T) {
	tx := db.Begin()
	dao := &AccountDao{
		db:     tx,
		logger: logger,
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
		tx.Rollback()
	}()
	Convey("更新账户余额", t, func() {
		a := &Account{
			AccountNo:   kd.Next().String(),
			AccountName: "测试资金账户",
			UserId:      kd.Next().String(),
			Username:    sql.NullString{String: "测试用户", Valid: true},
			Balance:     decimal.NewFromFloat(100),
			Status:      1,
			AccountType: 2,
		}
		id, err := dao.Insert(a)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
		//1. 增加余额
		Convey("增加余额", func() {
			r, err := dao.UpdateBalance(a.AccountNo, decimal.NewFromFloat(100))
			So(err, ShouldBeNil)
			So(r, ShouldEqual, 1)
			na := dao.GetByAccountNo(a.AccountNo)
			So(na, ShouldNotBeNil)
			newBalance := a.Balance.Add(decimal.NewFromFloat(100))
			So(na.Balance.String(), ShouldEqual, newBalance.String())
			//2. 扣减余额，余额足够
			Convey("扣减余额，余额充足", func() {
				amount := decimal.NewFromFloat(-20)
				r, err := dao.UpdateBalance(a.AccountNo, amount)
				So(err, ShouldBeNil)
				So(r, ShouldEqual, 1)
				na := dao.GetByAccountNo(a.AccountNo)
				So(na, ShouldNotBeNil)
				newBalance = newBalance.Add(amount)
				So(na.Balance.String(), ShouldEqual, newBalance.String())

				//3. 扣减余额，余额不足
				Convey("扣减余额，余额不足", func() {
					amount = decimal.NewFromFloat(-200)
					r, err := dao.UpdateBalance(a.AccountNo, amount)
					So(err, ShouldBeNil)
					So(r, ShouldEqual, 0)
					na := dao.GetByAccountNo(a.AccountNo)
					So(na, ShouldNotBeNil)
					So(na.Balance.String(), ShouldEqual, newBalance.String())
				})
			})
		})

	})

}
