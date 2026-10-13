package mods

import "github.com/ez4bk/ezboot/tools/reflect"

type Mod interface {
	Name() string
	Target() string
	DaoTarget() string
	Policy() *reflect.Policy
}
