package gorilla_test

import (
	"testing"

	"github.com/driver005/commerce/session/gorilla"
	"github.com/driver005/commerce/session/test"
	"github.com/gorilla/sessions"
)

func TestAll(t *testing.T) {
	engine := sessions.NewCookieStore([]byte("something-very-secret"))
	manager := gorilla.New("_session", engine)
	test.TestAll(manager, t)
}
