package v1

import (
	"blog_server/utils"

	"github.com/gin-gonic/gin"
)

type Category struct {
}

func (c *Category) Get(ctx *gin.Context) {
	utils.ResultOK(ctx, []string{"aaa", "bbb", "ccc"})
}

func (c *Category) GetByName(ctx *gin.Context) {
	cate_name := ctx.Param("catename")
	utils.ResultOK(ctx, []string{cate_name, "ddddd", "bbbbb"})
}
