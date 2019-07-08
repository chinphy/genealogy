package model

import "github.com/jinzhu/gorm"

// Clan 族谱模型
type Clan struct {
	gorm.Model
	Name   string
	Status string
}
