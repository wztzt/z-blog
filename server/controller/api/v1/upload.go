package v1

import (
	"blog_server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func (u *Upload) Post(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	ctx.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))
	utils.ResultOK(ctx, "ok")
	//src, _ := file.Open()
	//defer src.Close()
	//data, _ := io.ReadAll(src)
	//base64.StdEncoding.EncodeToString(data)

}
