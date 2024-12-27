package tools

import "github.com/gin-gonic/gin"

func GetInstanceName(c *gin.Context) string {
	return c.GetString("gin-ui-instance-name")
}
