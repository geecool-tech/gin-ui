package main

import (
	"github.com/geecool-tech/gin-ui/app"
	"github.com/geecool-tech/gin-ui/observer"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

func main() {
	engine := gin.Default()

	engine.Use(observer.New(
		"测试实例",
		app.WithAutoOpenConsole(true),
		app.WithPort(9981),
	),
	)
	engine.GET("/test", func(context *gin.Context) {
		time.Sleep(2 * time.Second)
	})
	engine.POST("/post", func(context *gin.Context) {
		var db *gorm.DB
		var data struct {
			Name string `json:"name"`
		}
		db.Find(&data)
		context.JSON(200, gin.H{
			"data": data,
		})
	})
	engine.Run(":9999")
}
