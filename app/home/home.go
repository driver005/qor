package home

import (
	"github.com/driver005/commerce/config/application"
	"github.com/driver005/commerce/render"
	"github.com/driver005/commerce/utils/funcmapmaker"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("home")}, "app/home/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	application.Router.Get("/", controller.Index)
	application.Router.Get("/switch_locale", controller.SwitchLocale)
}
