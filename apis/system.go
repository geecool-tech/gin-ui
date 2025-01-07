package apis

import (
	"fmt"
	"github.com/geecool-tech/gin-ui/tools"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/mem"
	"os"
)

type SystemController struct {
}

func (s *SystemController) Info(c *gin.Context) {
	hostName, _ := os.Hostname()
	vMem, _ := mem.VirtualMemory()
	tools.SuccessResp(c, map[string]any{
		"os_name": hostName,
		"mem": map[string]any{
			"total": fmt.Sprintf("%.2f GiB", float64(vMem.Total)/1024/1024/1024),
			"used":  fmt.Sprintf("%.2f GiB", float64(vMem.Used)/1024/1024/1024),
			"free":  fmt.Sprintf("%.2f GiB", float64(vMem.Available)/1024/1024/1024),
		},
	})
}
