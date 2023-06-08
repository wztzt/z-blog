package routes

import (
	v1 "blog_server/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	apiv1 := r.Group("api/v1")
	{
		health := &v1.Health{}
		health.InitRouter(apiv1)
		user := &v1.User{}
		user.InitRouter(apiv1)
		systeminfo := &v1.SystemInfo{}
		systeminfo.InitRouter(apiv1)
		command := &v1.Command{}
		apiv1.GET("cmd", command.Get)
	}
	return r
}
