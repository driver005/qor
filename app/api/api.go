package api

import (
	"github.com/driver005/commerce/admin"
	"github.com/driver005/commerce/config/application"
	"github.com/driver005/commerce/config/db"
	"github.com/driver005/commerce/models/orders"
	"github.com/driver005/commerce/models/products"
	"github.com/driver005/commerce/models/users"
	"github.com/driver005/commerce/qor"
)

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/api"
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
	API := admin.New(&qor.Config{DB: db.DB})

	Product := API.AddResource(&products.Product{})

	ColorVariationMeta := Product.Meta(&admin.Meta{Name: "ColorVariations"})
	ColorVariation := ColorVariationMeta.Resource
	ColorVariation.IndexAttrs("ID", "Color", "Images", "SizeVariations")
	ColorVariation.ShowAttrs("Color", "Images", "SizeVariations")

	SizeVariationMeta := ColorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	SizeVariation := SizeVariationMeta.Resource
	SizeVariation.IndexAttrs("ID", "Size", "AvailableQuantity")
	SizeVariation.ShowAttrs("ID", "Size", "AvailableQuantity")

	API.AddResource(&orders.Order{})

	API.AddResource(&users.User{})
	// User := API.AddResource(&users.User{})
	// userOrders, _ := User.AddSubResource("Orders")
	// userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})

	API.AddResource(&products.Category{})

	application.Router.Mount(app.Config.Prefix, API.NewServeMux(app.Config.Prefix))
}
