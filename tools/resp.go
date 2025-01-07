package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResp(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "success",
		"data": data,
	})
}
