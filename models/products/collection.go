package products

import (
	"github.com/driver005/commerce/l10n"
	"github.com/jinzhu/gorm"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
}
