package main

import (
	"github.com/tietang/props/v3/ini"
	"red-packet/infra"
	_ "red-packet/starters"
)

func main() {
	conf := ini.NewIniFileCompositeConfigSource("config.ini")
	app := infra.New(conf)
	app.Start()
}
