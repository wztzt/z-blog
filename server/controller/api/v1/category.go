package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Category struct {
}

func (c *Category) Get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []string{"aaa", "bbb", "ccc", "ddd"})
}
