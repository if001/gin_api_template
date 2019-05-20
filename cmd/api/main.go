// mainHTTPSample.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin_api_template/interfaces/api/handler"
	"gin_api_template/interfaces/api/middleware"
	"gin_api_template/interfaces/api/router"
)


func main() {
	var addr =  ":8080"

	handlers := handler.NewApiHandler()

	fmt.Printf("[START] server. port: %s\n", addr)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.SampleMiddleware())

	router.Routes(r, handlers)
	if err := r.Run(); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}