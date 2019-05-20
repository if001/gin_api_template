package router

import (
	"github.com/gin-gonic/gin"
	"gin_api_template/interfaces/api/handler"
	"gin_api_template/interfaces/api/middleware"
)


func Route(r *gin.Engine, h handler.ApiHandler) {

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/books", h.GetBooks)
	}

}
