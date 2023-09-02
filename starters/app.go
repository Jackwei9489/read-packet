package starters

import (
	"red-packet/infra"
	"red-packet/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
}
