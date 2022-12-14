package admin

import (
	"github.com/driver005/commerce/action_bar"
	"github.com/driver005/commerce/admin"
	"github.com/driver005/commerce/config/application"
	"github.com/driver005/commerce/config/i18n"
	"github.com/driver005/commerce/help"
	"github.com/driver005/commerce/media/asset_manager"
	"github.com/driver005/commerce/media/media_library"
	"github.com/driver005/commerce/models/settings"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

// AssetManager asset manager
var AssetManager *admin.Resource

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/admin"
	}
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	Admin := application.Admin

	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// Add Translations
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: -1})

	// Add Setting
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Shop Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	SetupSEO(Admin)
	SetupWidget(Admin)
	SetupDashboard(Admin)

	application.Router.Mount(app.Config.Prefix, Admin.NewServeMux(app.Config.Prefix))
}
