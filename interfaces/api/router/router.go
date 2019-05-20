package router

import (
	"github.com/gin-gonic/gin"
	"gin_api_template/interfaces/api/handler"
)


func Route(r *gin.Engine, h handler.ApiHandler) {
	r.GET("/books", h.GetBooks)
}
