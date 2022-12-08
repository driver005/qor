package blogs

import (
	"github.com/driver005/commerce/page_builder"
	"github.com/driver005/commerce/publish2"
)

type Page struct {
	page_builder.Page

	publish2.Version
	publish2.Schedule
	publish2.Visible
}
