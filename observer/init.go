package observer

import (
	"fmt"
	"github.com/geecool-tech/gin-ui/apis"
	"github.com/geecool-tech/gin-ui/app"
	"github.com/geecool-tech/gin-ui/handler"
	"github.com/geecool-tech/gin-ui/tools"
	"github.com/gin-gonic/gin"

	"sync"
)

var once sync.Once

func New(instanceName string, options ...app.Option) gin.HandlerFunc {
	instance := &app.GinUI{
		Name:            instanceName,
		AutoOpenConsole: true,
		Port:            9920,
		DbPath:          "",
	}
	for _, option := range options {
		option(instance)
	}
	fmt.Println("****************** GIN UI ******************")
	fmt.Println("             Welcome to GIN-UI!")
	fmt.Println(fmt.Sprintf("[GIN-UI] [INFO] Instance %s has been registerd...\n", instance.Name))
	fmt.Println("[GIN-UI] [INFO] GIN-UI started...")
	if !instance.AutoOpenConsole {
		fmt.Println("[GIN-UI] [WARNING] You've turned off the auto-open console.")
		fmt.Println("[GIN-UI] [WARNING] Go to http://localhost:9920 to open the console.")
	} else {
		fmt.Println("[GIN-UI] [INFO] The console is about to open automatically.")
	}
	fmt.Println("****************** GIN UI ******************")
	go func() {
		once.Do(func() {
			initialize(instance)
		})
	}()
	return func(context *gin.Context) {
		handler.MainHandlerFunc(context, instance)
	}
}
func initialize(instance *app.GinUI) {
	engine := gin.New()
	gin.SetMode(gin.ReleaseMode)
	engine.Use(gin.Recovery()) // 仅启用恢复中间件，防止 panic
	apis.InitApis(engine, instance)
	db := app.GetDb()
	_ = db.AutoMigrate(&app.ApiRecord{})
	go func() {
		if instance.AutoOpenConsole {
			_ = tools.OpenBrowser(fmt.Sprintf("http://localhost:%d/console", instance.Port))
		}
	}()
	_ = engine.Run(fmt.Sprintf(":%d", instance.Port))

}
