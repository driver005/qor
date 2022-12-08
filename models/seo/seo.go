package seo

import (
	"github.com/driver005/commerce/l10n"
	"github.com/driver005/commerce/seo"
)

type MySEOSetting struct {
	seo.QorSEOSetting
	l10n.Locale
}

type SEOGlobalSetting struct {
	SiteName string
}

var SEOCollection *seo.Collection
