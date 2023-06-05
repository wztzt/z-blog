package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func (h *Health) InitRouter(group *gin.RouterGroup) {
	group.GET("health", h.Check)
}

func (h *Health) Check(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{})
}
