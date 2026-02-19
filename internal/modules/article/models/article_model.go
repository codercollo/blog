package models

import (
	"github.com/codercollo/blog/internal/modules/user/models"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"type:varchar(191)"`
	Content string `gorm:"type:text"`
	UserID  uint
	User    models.User
}
