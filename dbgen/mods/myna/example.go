package myna

import (
	"time"

	"github.com/ez4bk/ezboot/snowflake"
	"github.com/ez4bk/ezboot/tools/reflect"
)

var Example = &reflect.TablePolicy{
	Package: "example.ezboot/myna/internal/dao/models",
	Columns: map[string]*reflect.ColumnPolicy{
		"id": {
			TypeValue: snowflake.Zero,
		},
		"type": {
			TypeValue: uint8(0),
		},
	},
	DAOPolicy: &reflect.DAOPolicy{
		Constraint: map[string]*reflect.DAOListConstraint{
			"name": {Likely: true},
			"type": {OptionalSkip: uint8(0)},
		},
	},
	PagerPolicy: &reflect.PagerPolicy{Type: reflect.PagerTypeToken},
	CachePolicy: &reflect.CachePolicy{
		Prefix:  "myna:db:example:",
		Expired: time.Hour * 24 * 7,
	},
}
