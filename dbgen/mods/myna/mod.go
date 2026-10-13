package myna

import (
	"github.com/ez4bk/ezboot/tools/reflect"
)

type Mod struct{}

func (m *Mod) Name() string {
	return "myna"
}

func (m *Mod) Target() string {
	return "internal/dao/models"
}

func (m *Mod) DaoTarget() string {
	return "internal/dao"
}

func (m *Mod) Policy() *reflect.Policy {
	return &reflect.Policy{
		Tables: map[string]*reflect.TablePolicy{
			"example": Example,
		},
	}
}
