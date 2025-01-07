package handler

import (
	"bytes"
	"fmt"
	"github.com/geecool-tech/gin-ui/app"
	"github.com/geecool-tech/gin-ui/tools"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func MainHandlerFunc(c *gin.Context, instance *app.GinUI) {
	// 提前读取请求体
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to read request body"})
		return
	}
	bodyStr := string(body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	start := time.Now()
	var errMsg string
	defer func() {
		if r := recover(); r != nil {
			errMsg = fmt.Sprintf("panic: %v", r)
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		headers := c.Request.Header
		var headerStr strings.Builder
		for key, values := range headers {
			headerStr.WriteString(fmt.Sprintf("%s: %s\n", key, strings.Join(values, ", ")))
		}
		status := c.Writer.Status()
		var duration time.Duration
		duration = time.Since(start)
		// 将记录写入数据库
		app.GetDb().Create(&app.ApiRecord{
			InstanceName: tools.GetInstanceName(c),
			Method:       c.Request.Method,
			Path:         c.Request.URL.Path,
			QueryParam:   c.Request.URL.RawQuery,
			Status:       status,
			Duration:     duration.Milliseconds(),
			IP:           c.ClientIP(),
			UserAgent:    c.Request.UserAgent(),
			Body:         bodyStr,
			Headers:      headerStr.String(),
			ErrorMsg:     errMsg,
		})
	}()

	// 执行后续中间件或处理函数
	c.Next()

}
