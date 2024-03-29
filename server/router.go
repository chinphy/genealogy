package server

import (
	"genealogy/api"
	"genealogy/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	r.GET("/ws", api.WsConnection)

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		needAuth := v1.Group("/")
		needAuth.Use(middleware.AuthRequired())
		{
			// User Routing
			needAuth.GET("user/me", api.UserMe)
			needAuth.DELETE("user/logout", api.UserLogout)
		}

		v1.POST("member/create", api.CreateClanMember)
	}
	return r
}
