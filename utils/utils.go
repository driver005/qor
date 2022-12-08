package utils

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/driver005/commerce/config/auth"
	"github.com/driver005/commerce/config/db"
	"github.com/driver005/commerce/l10n"
	"github.com/driver005/commerce/models/users"
	"github.com/driver005/commerce/qor/utils"
	"github.com/driver005/commerce/session"
	"github.com/driver005/commerce/session/manager"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	"github.com/microcosm-cc/bluemonday"
)

// GetCurrentUser get current user from request
func GetCurrentUser(req *http.Request) *users.User {
	if currentUser, ok := auth.Auth.GetCurrentUser(req).(*users.User); ok {
		return currentUser
	}
	return nil
}

// GetCurrentLocale get current locale from request
func GetCurrentLocale(req *http.Request) string {
	locale := l10n.Global
	if cookie, err := req.Cookie("locale"); err == nil {
		locale = cookie.Value
	}
	return locale
}

// GetDB get DB from request
func GetDB(req *http.Request) *gorm.DB {
	if db := utils.GetDBFromRequest(req); db != nil {
		return db
	}
	return db.DB
}

// URLParam get url params from request
func URLParam(name string, req *http.Request) string {
	return chi.URLParam(req, name)
}

// AddFlashMessage helper
func AddFlashMessage(w http.ResponseWriter, req *http.Request, message string, mtype string) error {
	return manager.SessionManager.Flash(w, req, session.Message{Message: template.HTML(message), Type: mtype})
}

// HTMLSanitizer HTML sanitizer
var HTMLSanitizer = bluemonday.UGCPolicy()

func FormatPrice(price interface{}) string {
	switch price.(type) {
	case float32, float64:
		return fmt.Sprintf("%0.2f", price)
	case int, uint, int32, int64, uint32, uint64:
		return fmt.Sprintf("%d.00", price)
	}
	return ""
}
