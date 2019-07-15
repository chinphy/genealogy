package api

import (
	"genealogy/serializer"
	"genealogy/service"

	"github.com/gin-gonic/gin"
)

// CreateClanMember 创建族员
func CreateClanMember(c *gin.Context) {
	var service service.ClanMemberService

	if err := c.ShouldBind(&service); err == nil {
		if member, err := service.CreateMember(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildClanMemberResponse(member)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
