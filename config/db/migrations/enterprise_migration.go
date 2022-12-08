//go:build enterprise
// +build enterprise

package migrations

import "github.com/driver005/commerce/app/enterprise"

func init() {
	AutoMigrate(&enterprise.QorMicroSite{})
}
