package cmd

import (
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/driver005/commerce/admin"
	"github.com/driver005/commerce/app/account"
	adminapp "github.com/driver005/commerce/app/admin"
	"github.com/driver005/commerce/app/api"
	"github.com/driver005/commerce/app/enterprise"
	"github.com/driver005/commerce/app/home"
	"github.com/driver005/commerce/app/orders"
	"github.com/driver005/commerce/app/pages"
	"github.com/driver005/commerce/app/products"
	"github.com/driver005/commerce/app/static"
	"github.com/driver005/commerce/config"
	"github.com/driver005/commerce/config/application"
	"github.com/driver005/commerce/config/auth"
	"github.com/driver005/commerce/config/bindatafs"
	"github.com/driver005/commerce/config/db"
	"github.com/driver005/commerce/publish2"
	"github.com/driver005/commerce/qor"
	"github.com/driver005/commerce/qor/utils"
	"github.com/driver005/commerce/utils/funcmapmaker"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/spf13/cobra"
)

var serveControls = `## Configuration
Ory Hydra can be configured using environment variables as well as a configuration file. For more information
on configuration options, open the configuration documentation:
>> https://www.ory.sh/hydra/docs/reference/configuration <<
`

func NewServeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Parent command for starting public and administrative HTTP/2 APIs",
		Long: `Ory Hydra exposes two ports, a public and an administrative port. The public port is responsible
		for handling requests from the public internet, such as the OAuth 2.0 Authorize and Token URLs. The administrative
		port handles administrative requests like creating OAuth 2.0 Clients, managing JSON Web Keys, and managing User Login
		and Consent sessions.
		It is recommended to run "hydra serve all". If you need granular control over CORS settings or similar, you may
		want to run "hydra serve admin" and "hydra serve public" separately.
		To learn more about each individual command, run:
		- hydra help serve all
		- hydra help serve admin
		- hydra help serve public
		All sub-commands share command line flags and configuration options.
		` + serveControls,
	}

	return cmd
}

func NewServeAllCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "Serves both public and administrative HTTP/2 APIs",
		Long: `Starts a process which listens on two ports for public and administrative HTTP/2 API requests.
If you want more granular control (e.g. different TLS settings) over each API group (administrative, public) you
can run "serve admin" and "serve public" separately.
This command exposes a variety of controls via environment variables. You can
set environments using "export KEY=VALUE" (Linux/macOS) or "set KEY=VALUE" (Windows). On Linux,
you can also set environments by prepending key value pairs: "KEY=VALUE KEY2=VALUE2 hydra"
All possible controls are listed below. This command exposes exposes command line flags, which are listed below
the controls section.
` + serveControls,
		Run: func(cmd *cobra.Command, args []string) {
			RunServeAll()
		},
	}
}

func RunServeAll() {
	var (
		Router = chi.NewRouter()
		Admin  = admin.New(&admin.AdminConfig{
			SiteName: "QOR DEMO",
			Auth:     auth.AdminAuth{},
			DB:       db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
		})
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// for demo, don't use this for your production site
			w.Header().Add("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.Header.Del("Authorization")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var (
				tx         = db.DB
				qorContext = &qor.Context{Request: req, Writer: w}
			)

			if locale := utils.GetLocale(qorContext); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			ctx := context.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	})

	Application.Use(api.New(&api.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(products.New(&products.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(orders.New(&orders.Config{}))
	Application.Use(pages.New(&pages.Config{}))
	Application.Use(enterprise.New(&enterprise.Config{}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/system"},
		Handler: utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))),
	}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"},
		Handler: bindatafs.AssetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"),
	}))

	fmt.Printf("Listening on: %v\n", config.Config.Port)
	if config.Config.HTTPS {
		if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Config.Port), "config/local_certs/server.crt", "config/local_certs/server.key", Application.NewServeMux()); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), Application.NewServeMux()); err != nil {
			panic(err)
		}
	}
}
