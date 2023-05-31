package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Add(ctx *gin.Context) {

}

func (u *User) Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "login")
}

func (u *User) Options(ctx *gin.Context) {

}
