package model

import "github.com/jinzhu/gorm"

// ClanMember 族员模型
type ClanMember struct {
	gorm.Model
	UserID uint `gorm:"default:0"`
	Name   string
	sex    string
	Avatar string `gorm:"size:1000"`
	Status string
}
