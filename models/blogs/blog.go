package blogs

import (
	"github.com/driver005/commerce/models/users"
	"github.com/driver005/commerce/publish2"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Author   users.User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
	publish2.Version
	publish2.Schedule
	publish2.Visible
}
