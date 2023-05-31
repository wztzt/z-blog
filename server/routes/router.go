package routes

import (
	v1 "blog_server/controller/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("api/v1")
	{
		user := v1.User{}
		apiv1.POST("login", user.Login)
		apiv1.GET("login", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "login --")
		})
		apiv1.OPTIONS("login", user.Options)
		systeminfo := v1.SystemInfo{}
		apiv1.POST("systeminfo", systeminfo.Post)
	}
	return r
}
