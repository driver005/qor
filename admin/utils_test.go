package admin

import (
	"fmt"
	"testing"
)

func TestGetDepVersionFromMod(t *testing.T) {
	goModDeps = []string{
		"github.com/driver005/commerce/l10n v0.0.0-20181031091737-2ca95fb3b4dd // indirect",
		"github.com/driver005/commerce/admin v1.1.1",
		"github.com/driver005/commerce/publish2 v1.1.0 // indirect",
		"github.com/driver005/commerce/media v0.0.0-20191022071353-19cf289e17d4 // indirect",
		"github.com/driver005/commerce/i18n v2.0.7",
	}
	cases := []struct {
		view string
		want string
	}{
		{view: "github.com/driver005/commerce/l10n/views", want: "pkg/mod/github.com/driver005/commerce/l10n@v0.0.0-20181031091737-2ca95fb3b4dd/views"},
		{view: "github.com/driver005/commerce/admin/views", want: "pkg/mod/github.com/driver005/commerce/admin@v1.1.1/views"},
		{view: "github.com/driver005/commerce/publish2/views", want: "pkg/mod/github.com/driver005/commerce/publish2@v1.1.0/views"},
		{view: "github.com/driver005/commerce/media/media_library/views", want: "pkg/mod/github.com/driver005/commerce/media@v0.0.0-20191022071353-19cf289e17d4/media_library/views"},
		{view: "github.com/driver005/commerce/i18n/exchange_actions/views", want: "pkg/mod/github.com/driver005/commerce/i18n@v2.0.7/exchange_actions/views"},
		{view: "no/unknown/nonexistent", want: "no/unknown/nonexistent"},
	}
	for _, v := range cases {
		if got := getDepVersionFromMod(v.view); v.want != got {
			t.Errorf("GetDepVersionFromMod-viewpath: %v, want: %v, got: %v", v.view, v.want, got)
		} else {
			fmt.Println(got)
		}
	}
}
