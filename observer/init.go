package observer

import (
	"fmt"
	"github.com/geecool-tech/gin-ui/apis"
	"github.com/geecool-tech/gin-ui/handler"
	"github.com/geecool-tech/gin-ui/tools"
	"github.com/gin-gonic/gin"
	"sync"
)

var once sync.Once
var GlobalInstanceName string
var OpenConsole bool

func New(instanceName string, openConsole bool) gin.HandlerFunc {
	fmt.Println("****************** GIN UI ******************")
	fmt.Println("             Welcome to GIN-UI!")
	fmt.Println("[GIN-UI] [INFO] GIN-UI started...")
	if !openConsole {
		fmt.Println("[GIN-UI] [WARNING] You've turned off the auto-open console.")
		fmt.Println("[GIN-UI] [WARNING] Go to http://localhost:9920 to open the console.")
	} else {
		fmt.Println("[GIN-UI] [INFO] The console is about to open automatically.")
	}
	fmt.Println("****************** GIN UI ******************")
	GlobalInstanceName = instanceName
	OpenConsole = openConsole
	go func() {
		once.Do(initialize)
	}()
	return func(context *gin.Context) {
		context.Set("gin-ui-instance-name", GlobalInstanceName)
		context.Set("gin-ui-open-console", openConsole)
		handler.MainHandlerFunc(context)
	}
}
func initialize() {
	engine := gin.New()
	gin.SetMode(gin.ReleaseMode)
	engine.Use(gin.Recovery()) // 仅启用恢复中间件，防止 panic
	apis.InitApis(engine)
	go func() {
		if OpenConsole {
			_ = tools.OpenBrowser("http://localhost:9920/console")
		}
	}()
	_ = engine.Run(":9920")

}
