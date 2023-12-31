package starters

import (
	"red-packet/infra"
	"red-packet/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.LogStarter{})
	infra.Register(&base.GormStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GinServerStarter{})
}
