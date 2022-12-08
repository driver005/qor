package manager

import (
	"net/http"

	"github.com/driver005/commerce/middlewares"
	"github.com/driver005/commerce/session"
	"github.com/driver005/commerce/session/gorilla"
	"github.com/gorilla/sessions"
)

// SessionManager default session manager
var SessionManager session.ManagerInterface = gorilla.New("_session", sessions.NewCookieStore([]byte("secret")))

func init() {
	middlewares.Use(middlewares.Middleware{
		Name: "session",
		Handler: func(handler http.Handler) http.Handler {
			return SessionManager.Middleware(handler)
		},
	})
}
