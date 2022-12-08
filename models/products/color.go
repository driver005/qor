package products

import (
	"strings"

	"github.com/driver005/commerce/l10n"
	"github.com/driver005/commerce/publish2"
	"github.com/driver005/commerce/sorting"
	"github.com/driver005/commerce/validations"
	"github.com/jinzhu/gorm"
)

type Color struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`

	publish2.Version
	publish2.Schedule
	publish2.Visible
}

func (color Color) Validate(db *gorm.DB) {
	if strings.TrimSpace(color.Name) == "" {
		db.AddError(validations.NewError(color, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(color.Code) == "" {
		db.AddError(validations.NewError(color, "Code", "Code can not be empty"))
	}
}
