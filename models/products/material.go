package products

import (
	"github.com/driver005/commerce/l10n"
	"github.com/jinzhu/gorm"
)

type Material struct {
	gorm.Model
	l10n.Locale
	Name string
	Code string `l10n:"sync"`
}
