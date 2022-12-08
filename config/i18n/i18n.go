package i18n

import (
	"path/filepath"

	"github.com/driver005/commerce/i18n"
	"github.com/driver005/commerce/i18n/backends/database"
	"github.com/driver005/commerce/i18n/backends/yaml"

	"github.com/driver005/commerce/config"
	"github.com/driver005/commerce/config/db"
)

var I18n *i18n.I18n

func init() {
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
}
