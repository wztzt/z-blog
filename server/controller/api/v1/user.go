package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) InitRouter(group *gin.RouterGroup) {
	group.POST("login", u.Login)
}

func (u *User) Add(ctx *gin.Context) {

}

func (u *User) Login(ctx *gin.Context) {
	fmt.Println(ctx.Request.Header)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "login",
		"host_name": ctx.Request.Host,
	})
}
