package apis

import (
	"github.com/geecool-tech/gin-ui/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitApis(e *gin.Engine, instance *app.GinUI) {
	e.GET("/console", func(context *gin.Context) {
		context.String(http.StatusOK, instance.Name)
	})
	{
		systemRouter := e.Group("/system")
		var systemController SystemController
		systemRouter.GET("/info", systemController.Info)
	}
}
