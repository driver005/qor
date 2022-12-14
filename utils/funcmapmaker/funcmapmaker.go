package funcmapmaker

import (
	"html/template"
	"net/http"

	"github.com/driver005/commerce/action_bar"
	"github.com/driver005/commerce/app/admin"
	"github.com/driver005/commerce/config"
	"github.com/driver005/commerce/config/i18n"
	"github.com/driver005/commerce/i18n/inline_edit"
	"github.com/driver005/commerce/models/products"
	"github.com/driver005/commerce/models/seo"
	"github.com/driver005/commerce/models/users"
	"github.com/driver005/commerce/qor"
	"github.com/driver005/commerce/render"
	"github.com/driver005/commerce/session"
	"github.com/driver005/commerce/session/manager"
	"github.com/driver005/commerce/utils"
	"github.com/driver005/commerce/widget"
)

// GetEditMode get edit mode
func GetEditMode(w http.ResponseWriter, req *http.Request) bool {
	return admin.ActionBar.EditMode(w, req)
}

// AddFuncMapMaker add FuncMapMaker to view
func AddFuncMapMaker(view *render.Render) *render.Render {
	oldFuncMapMaker := view.FuncMapMaker
	view.FuncMapMaker = func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
		funcMap := template.FuncMap{}
		if oldFuncMapMaker != nil {
			funcMap = oldFuncMapMaker(render, req, w)
		}

		// Add `t` method
		for key, fc := range inline_edit.FuncMap(i18n.I18n, utils.GetCurrentLocale(req), GetEditMode(w, req)) {
			funcMap[key] = fc
		}

		for key, value := range admin.ActionBar.FuncMap(w, req) {
			funcMap[key] = value
		}

		widgetContext := admin.Widgets.NewContext(&widget.Context{
			DB:         utils.GetDB(req),
			Options:    map[string]interface{}{"Request": req},
			InlineEdit: GetEditMode(w, req),
		})
		for key, fc := range widgetContext.FuncMap() {
			funcMap[key] = fc
		}

		funcMap["raw"] = func(str string) template.HTML {
			return template.HTML(utils.HTMLSanitizer.Sanitize(str))
		}

		funcMap["flashes"] = func() []session.Message {
			return manager.SessionManager.Flashes(w, req)
		}

		// Add `action_bar` method
		funcMap["render_action_bar"] = func() template.HTML {
			return admin.ActionBar.Actions(action_bar.Action{Name: "Edit SEO", Link: seo.SEOCollection.SEOSettingURL("/help")}).Render(w, req)
		}

		funcMap["render_seo_tag"] = func() template.HTML {
			return seo.SEOCollection.Render(&qor.Context{DB: utils.GetDB(req)}, "Default Page")
		}

		funcMap["get_categories"] = func() (categories []products.Category) {
			utils.GetDB(req).Find(&categories)
			return
		}

		funcMap["current_locale"] = func() string {
			return utils.GetCurrentLocale(req)
		}

		funcMap["current_user"] = func() *users.User {
			return utils.GetCurrentUser(req)
		}

		funcMap["related_products"] = func(cv products.ColorVariation) []products.Product {
			var products []products.Product
			utils.GetDB(req).Preload("ColorVariations").Limit(4).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["other_also_bought"] = func(cv products.ColorVariation) []products.Product {
			var products []products.Product
			utils.GetDB(req).Preload("ColorVariations").Order("id ASC").Limit(8).Find(&products, "id <> ?", cv.ProductID)
			return products
		}

		funcMap["amazon_payment_gateway"] = func() interface{} {
			return config.Config.AmazonPay
		}

		funcMap["format_price"] = func(price interface{}) string {
			return utils.FormatPrice(price)
		}

		return funcMap
	}

	return view
}
