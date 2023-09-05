package base

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"red-packet/infra"
)

const (
	defaultDbUrl = "root:wz7953622@tcp(127.0.0.1:3306)/red_packet?charset=utf8mb4&parseTime=True&loc=Local"
)

var database *gorm.DB

func DataBase() *gorm.DB {
	return database
}

type GormStarter struct {
	infra.BaseStarter
}

func (o *GormStarter) SetUp(ctx infra.StarterContext) {
	dsn := ctx.Props().GetDefault("database.mysql.connect.url", defaultDbUrl)

	//logger
	//logger := zapgorm2.New(logger.Desugar())
	//logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks
	//logger.LogMode(logger2.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger2.Default.LogMode(logger2.Info)})
	log := Logger()
	if err != nil {
		panic(err)
	}
	log.Info("database setup success...")
	database = db
	//migrate
	//_ = db.AutoMigrate(&accounts.Account{})
}
