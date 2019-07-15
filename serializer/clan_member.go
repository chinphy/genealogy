package serializer

import "genealogy/model"

// ClanMember 族员序列化器
type ClanMember struct {
	ID        uint   `json:"id"`
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}

// ClanMemberResponse 单个族员序列化
type ClanMemberResponse struct {
	Response
	Data ClanMember `json:"data"`
}

// BuildClanMember 序列化用户
func BuildClanMember(item model.ClanMember) ClanMember {
	return ClanMember{
		ID:        item.ID,
		Name:      item.Name,
		Status:    item.Status,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildClanMemberResponse 序列化族员响应
func BuildClanMemberResponse(item model.ClanMember) ClanMemberResponse {
	return ClanMemberResponse{
		Data: BuildClanMember(item),
	}
}
