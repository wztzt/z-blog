package v1

import (
	"blog_server/utils"
	"blog_server/utils/redis"
	"math/rand"
	"strings"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type User struct {
	User     string `json:"user"`
	PassWord string `json:"password"`
}

func (u *User) InitRouter(group *gin.RouterGroup) {
	group.POST("login", u.Login)
	group.POST("register", u.Register)
}

func (u *User) Add(ctx *gin.Context) {

}

func (u *User) Login(ctx *gin.Context) {
	user := User{}
	if err := ctx.ShouldBind(&user); err != nil {
		utils.ResultFail(ctx, utils.CODE_BODY_ERROR, err)
		return
	}
	token, err := ctx.Cookie("token")
	if err != nil {
		utils.ResultFail(ctx, utils.CODE_REDIS_ERROR, err)
		return
	}
	utils.ResultOK(ctx, gin.H{
		"token": token,
	})
}

func (u *User) Register(ctx *gin.Context) {
	user := User{}
	if err := ctx.ShouldBind(&user); err != nil {
		utils.ResultFail(ctx, utils.CODE_BODY_ERROR, err)
		return
	}
	//存储数据库
	token := u.gen_token()
	if err := redis.Set("user_"+user.User, token, 0).Err(); err != nil {
		utils.ResultFail(ctx, utils.CODE_REDIS_ERROR, err)
		return
	}
	ctx.SetCookie("token", token, 60*2, "/", "", true, false)
	utils.ResultOK(ctx, gin.H{})
}

func (u *User) Logout(ctx *gin.Context) {

}

func (u *User) gen_token() string {
	n := 16
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
