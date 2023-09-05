package accounts

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountLogDao struct {
	db     *gorm.DB
	logger *zap.SugaredLogger
}

// 通过流水编号查询流水记录
