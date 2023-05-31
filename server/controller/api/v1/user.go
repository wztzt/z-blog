package v1

import (
	"blog_server/utils"
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
	ctx.SetCookie("Token", "token_token_token", 24*60*60, "/", utils.Domain, true, false)
	ctx.JSON(http.StatusOK, gin.H{
		"msg":       "login",
		"host_name": ctx.Request.Host,
	})
}
