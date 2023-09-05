package base

import (
	"github.com/tietang/props/v3/kvs"
	"red-packet/infra"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (o *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
}
