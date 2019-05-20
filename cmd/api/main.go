// mainHTTPSample.go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gin_api_template/interfaces/api/handler"
	"gin_api_template/interfaces/api/router"
	"gin_api_template/interfaces/api/middleware"
)


func main() {
	var addr =  ":8080"

	handlers := handler.NewApiHandler()

	fmt.Printf("[START] server. port: %s\n", addr)

	r := gin.Default()
	r.Use(middleware.SampleMiddleware())

	router.Route(r, handlers)
	if err := r.Run(); err != nil {
		panic(fmt.Errorf("[FAILED] start sever. err: %v", err))
	}
}