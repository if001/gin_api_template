package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type User struct {
	Name     string
	Password string
}

func getUser(token string) User { // move to application layer
	user := User{
		Name:"hoge",
		Password:"huga",
	}
	return user
}
func AuthRequired(c *gin.Context) {
	token := c.Query("token")
	if token == "hogehoge" {
		fmt.Println("authorize ok!")
		user := getUser(token)
		c.Set(gin.AuthUserKey, user) // set login user
		return
	}
	c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
	c.Abort()
}
