package apis

import "github.com/gin-gonic/gin"

func InitApis(e *gin.Engine) {
	e.GET("/console")
}
