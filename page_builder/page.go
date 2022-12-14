package page_builder

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/driver005/commerce/admin"
	"github.com/driver005/commerce/qor"
	"github.com/driver005/commerce/qor/resource"
	"github.com/driver005/commerce/slug"
	"github.com/driver005/commerce/widget"
	"github.com/jinzhu/gorm"
)

type Page struct {
	gorm.Model
	Title         string
	TitleWithSlug slug.Slug

	Containers Containers `sql:"type:text;"`
}

func (*Page) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		res.UseTheme("page_builder")
	}
}

type Container struct {
	Name string
}

type Containers []Container

func (containers *Containers) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, containers)
	case string:
		if v != "" {
			return containers.Scan([]byte(v))
		}
	case []string:
		c := Containers{}
		for _, name := range v {
			if name != "" {
				c = append(c, Container{Name: name})
			}
		}
		*containers = c
	default:
		return errors.New("not supported")
	}

	return nil
}

func (containers Containers) Value() (driver.Value, error) {
	if len(containers) == 0 {
		return nil, nil
	}

	return json.Marshal(containers)
}

func (page Page) GetContainerRecords(db *gorm.DB) (records []widget.QorWidgetSetting) {
	names := []string{}
	for _, container := range page.Containers {
		names = append(names, container.Name)
	}

	containers := []widget.QorWidgetSetting{}
	db.Where("name in (?) AND scope = ?", names, "default").Find(&containers)

	for _, name := range names {
		for _, container := range containers {
			if container.Name == name {
				records = append(records, container)
			}
		}
	}

	return
}

type Config struct {
	Admin       *admin.Admin
	PageModel   interface{}
	Containers  *widget.Widgets
	AdminConfig *admin.Config
}

func New(config *Config) *admin.Resource {
	resource := config.Admin.AddResource(config.PageModel, config.AdminConfig)

	resource.Meta(&admin.Meta{
		Name: "Containers",
		Valuer: func(value interface{}, context *qor.Context) interface{} {
			nameWithIcons := [][]string{}
			if page, ok := value.(interface {
				GetContainerRecords(*gorm.DB) []widget.QorWidgetSetting
			}); ok {
				for _, container := range page.GetContainerRecords(context.GetDB()) {
					nameWithIcon := []string{container.Name}

					icon := widget.GetWidget(container.WidgetType).PreviewIcon
					nameWithIcon = append(nameWithIcon, icon)

					nameWithIcons = append(nameWithIcons, nameWithIcon)
				}
			}

			return nameWithIcons
		},
		Config: &admin.SelectManyConfig{
			SelectionTemplate:  "metas/form/sortable_widgets.tmpl",
			SelectMode:         "bottom_sheet",
			DefaultCreating:    true,
			RemoteDataResource: config.Containers.WidgetSettingResource,
		}})

	admin.RegisterViewPath("github.com/driver005/commerce/page_builder/views")
	return resource
}
