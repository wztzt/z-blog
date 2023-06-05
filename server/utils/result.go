package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CODE_SUCCESS      = 0
	CODE_PARAMS_ERROR = -1
	CODE_BODY_ERROR   = -2
	CODE_REDIS_ERROR  = -3
)

type result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResultOK(ctx *gin.Context, obj any) {
	ctx.JSON(http.StatusOK, result{
		Code:    CODE_SUCCESS,
		Message: "success",
		Data:    obj,
	})
}

func ResultFail(ctx *gin.Context, code int, err error) {
	ctx.JSON(http.StatusOK, result{
		Code:    code,
		Message: err.Error(),
	})
}
