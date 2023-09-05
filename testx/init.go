package testx

import (
	"github.com/tietang/props/v3/ini"
	"red-packet/infra"
	"red-packet/infra/base"
)

func init() {
	conf := ini.NewIniFileCompositeConfigSource("../../config.ini")
	app := infra.New(conf)

	infra.Register(&base.PropsStarter{})
	infra.Register(&base.LogStarter{})
	infra.Register(&base.GormStarter{})
	infra.Register(&base.ValidatorStarter{})

	app.Start()
}
