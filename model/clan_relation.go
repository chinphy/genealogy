package model

import "github.com/jinzhu/gorm"

// ClanMemberRelation 族员关系模型
type ClanMemberRelation struct {
	gorm.Model
	MemberID     uint `gorm:"default:0"`
	FatherID     uint `gorm:"default:0"`
	MotherID     uint `gorm:"default:0"`
	RelationType uint
	ClanID       uint `gorm:"default:0"`
	Status       string
}
