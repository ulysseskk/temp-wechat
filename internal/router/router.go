package router

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/url"
)

func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Any("/:path", func(context *gin.Context) {
		fmt.Println(fmt.Sprintf("Path %s", context.FullPath()))
		fmt.Println(fmt.Sprintf("Method %s", context.Request.Method))
		fmt.Println(fmt.Sprintf("URL %+v", context.Request.URL.String()))
		bodyBuffer := &bytes.Buffer{}
		_, err := bodyBuffer.ReadFrom(context.Request.Body)
		if err != nil {
			fmt.Println("Error ", err)
			return
		}
		fmt.Println(fmt.Sprintf("Body %+v", bodyBuffer.String()))
		fmt.Println(url.QueryUnescape(bodyBuffer.String()))
	})
	return engine
}
