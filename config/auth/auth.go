package auth

import (
	"time"

	"github.com/driver005/commerce/auth"
	"github.com/driver005/commerce/auth/authority"
	"github.com/driver005/commerce/auth/providers/facebook"
	"github.com/driver005/commerce/auth/providers/github"
	"github.com/driver005/commerce/auth/providers/google"
	"github.com/driver005/commerce/auth/providers/twitter"
	"github.com/driver005/commerce/auth_themes/clean"
	"github.com/driver005/commerce/config"
	"github.com/driver005/commerce/config/bindatafs"
	"github.com/driver005/commerce/config/db"
	"github.com/driver005/commerce/models/users"
	"github.com/driver005/commerce/render"
)

var (
	// Auth initialize Auth for Authentication
	Auth = clean.New(&auth.Config{
		DB:         db.DB,
		Mailer:     config.Mailer,
		Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("auth")}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority initialize Authority for Authorization
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {
	if config.Config.Github.ClientID != "" {
		Auth.RegisterProvider(github.New(&config.Config.Github))
	}
	if config.Config.Google.ClientID != "" {
		Auth.RegisterProvider(google.New(&config.Config.Google))
	}
	if config.Config.Facebook.ClientID != "" {
		Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	}
	if config.Config.Twitter.ClientID != "" {
		Auth.RegisterProvider(twitter.New(&config.Config.Twitter))
	}

	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
