package products

import (
	"fmt"
	"strings"

	"github.com/driver005/commerce/l10n"
	"github.com/driver005/commerce/sorting"
	"github.com/driver005/commerce/validations"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string

	Categories []Category
	CategoryID uint
}

func (category Category) Validate(db *gorm.DB) {
	if strings.TrimSpace(category.Name) == "" {
		db.AddError(validations.NewError(category, "Name", "Name can not be empty"))
	}
}

func (category Category) DefaultPath() string {
	if len(category.Code) > 0 {
		return fmt.Sprintf("/category/%s", category.Code)
	}
	return "/"
}
