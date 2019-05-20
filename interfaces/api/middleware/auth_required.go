package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthRequired(c *gin.Context) {
	token := c.Query("token")
	panic("auth panic")
	if token == "hogehoge" {
		c.JSON(http.StatusOK, http.StatusText(http.StatusOK))
		c.Next()
	}
	c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	c.Abort()
}

