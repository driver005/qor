package admin

import (
	"path"
	"regexp"
	"time"

	"github.com/driver005/commerce/admin"
	"github.com/driver005/commerce/config/db"
	"github.com/driver005/commerce/models/orders"
	"github.com/driver005/commerce/notification"
	"github.com/driver005/commerce/notification/channels/database"
)

// SetupNotification add notification
func SetupNotification(Admin *admin.Admin) {
	Notification := notification.New(&notification.Config{})
	Notification.RegisterChannel(database.New(&database.Config{DB: db.DB}))
	Notification.Action(&notification.Action{
		Name: "Confirm",
		Visible: func(data *notification.QorNotification, context *admin.Context) bool {
			return data.ResolvedAt == nil
		},
		MessageTypes: []string{"order_returned"},
		Handler: func(argument *notification.ActionArgument) error {
			orderID := regexp.MustCompile(`#(\d+)`).FindStringSubmatch(argument.Message.Body)[1]
			err := argument.Context.GetDB().Model(&orders.Order{}).Where("id = ? AND returned_at IS NULL", orderID).Update("returned_at", time.Now()).Error
			if err == nil {
				return argument.Context.GetDB().Model(argument.Message).Update("resolved_at", time.Now()).Error
			}
			return err
		},
		Undo: func(argument *notification.ActionArgument) error {
			orderID := regexp.MustCompile(`#(\d+)`).FindStringSubmatch(argument.Message.Body)[1]
			err := argument.Context.GetDB().Model(&orders.Order{}).Where("id = ? AND returned_at IS NOT NULL", orderID).Update("returned_at", nil).Error
			if err == nil {
				return argument.Context.GetDB().Model(argument.Message).Update("resolved_at", nil).Error
			}
			return err
		},
	})
	Notification.Action(&notification.Action{
		Name:         "Check it out",
		MessageTypes: []string{"order_paid_cancelled", "order_processed", "order_returned"},
		URL: func(data *notification.QorNotification, context *admin.Context) string {
			return path.Join("/admin/orders/", regexp.MustCompile(`#(\d+)`).FindStringSubmatch(data.Body)[1])
		},
	})
	Notification.Action(&notification.Action{
		Name:         "Dismiss",
		MessageTypes: []string{"order_paid_cancelled", "info", "order_processed", "order_returned"},
		Visible: func(data *notification.QorNotification, context *admin.Context) bool {
			return data.ResolvedAt == nil
		},
		Handler: func(argument *notification.ActionArgument) error {
			return argument.Context.GetDB().Model(argument.Message).Update("resolved_at", time.Now()).Error
		},
		Undo: func(argument *notification.ActionArgument) error {
			return argument.Context.GetDB().Model(argument.Message).Update("resolved_at", nil).Error
		},
	})
	Admin.NewResource(Notification)
}
