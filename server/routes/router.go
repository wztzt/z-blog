package routes

import (
	v1 "blog_server/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	apiv1 := r.Group("api/v1")
	{
		user := v1.User{}
		apiv1.POST("login", user.Login)
	}
	return r
}
