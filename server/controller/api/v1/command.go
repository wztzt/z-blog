package v1

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Command struct {
	Path string   `json:"path"`
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

func (c *Command) Get(ctx *gin.Context) {
	cmd := &Command{}
	cmd.Path = ctx.Query("path")
	if cmd.Path == "" {
		cmd.Path, _ = os.Getwd()
	}
	cmd.Cmd = ctx.Query("cmd")
	cmd.Args = ctx.QueryArray("args")
	execmd := exec.Command(cmd.Cmd, cmd.Args...)
	data, _ := execmd.Output()
	ctx.String(http.StatusOK, string(data))
}
