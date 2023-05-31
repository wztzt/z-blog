package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	CpuInfo  []cpu.InfoStat         `json:"cpu_info"`
	MemInfo  *mem.VirtualMemoryStat `json:"mem_info"`
	HostInfo *host.InfoStat         `json:"host_info"`
	UserInfo []host.UserStat        `json:"user_info"`
}

func (s *SystemInfo) Post(ctx *gin.Context) {
	info := &SystemInfo{}
	info.CpuInfo, _ = cpu.Info()
	info.MemInfo, _ = mem.VirtualMemory()
	info.HostInfo, _ = host.Info()
	info.UserInfo, _ = host.Users()
	ctx.JSON(http.StatusOK, info)

}
