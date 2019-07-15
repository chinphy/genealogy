package service

import (
	"genealogy/model"
	"genealogy/serializer"
)

// ClanMemberService 族员业务
type ClanMemberService struct {
	Name     string `json:"name" binding:"required,min=5,max=30"`
	Sex      string `json:"sex" binding:"required"`
	Avatar   string `json:"avatar"`
	FatherID uint   `json:"father_id"`
	MotherID uint   `json:"mother_id"`
	ClanID   uint   `json:"clan_id" binding:"required"`
}

// CreateMember 创建族员
func (srv *ClanMemberService) CreateMember() (model.ClanMember, *serializer.Response) {
	member := model.ClanMember{
		Name: srv.Name,
		Sex:  srv.Sex,
	}

	// 开启事务
	tx := model.DB.Begin()

	// 创建用户
	if err := tx.Create(&member).Error; err != nil {
		tx.Rollback()
		return member, &serializer.Response{
			Status: 40002,
			Msg:    "创建失败",
		}
	}

	relation := model.ClanMemberRelation{
		MemberID: member.ID,
		FatherID: srv.FatherID,
		MotherID: srv.MotherID,
		ClanID:   srv.ClanID,
	}
	// 创建关系
	if err := model.DB.Create(&relation).Error; err != nil {
		tx.Rollback()
		return member, &serializer.Response{
			Status: 40002,
			Msg:    "创建关系失败",
		}
	}
	tx.Commit()
	return member, nil
}
