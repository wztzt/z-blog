package v1

import (
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
	ctx.String(http.StatusOK, "login")
}
