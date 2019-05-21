package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"gin_api_template/interfaces/api/router"
	"gin_api_template/interfaces/api/handler"
	"strconv"
)

type URI struct {
	uri string
}

func (u *URI) toString() string {
	return u.uri
}
func (u *URI) withToken(token string) *URI {
	q := &QueryString{key:"token",value:token}
	u.uri = u.uri + u.LinkChar() + q.String()
	return u
}
func (u *URI) withQueryString(q QueryString) *URI {
	u.uri = u.uri + u.LinkChar() + q.String()
	return u
}
func (u *URI) LinkChar() string {
	linkChar := "?"
	if containQuestion(u.uri) {
		linkChar = "&"
	}
	return linkChar
}
func containQuestion(uri string) bool {
	uriRune := []rune(uri)
	for i := range uriRune {
		if string(uri[i]) == "?" {
			return true
		}
	}
	return false
}

type QueryString struct {
	key   string
	value interface{}
}

func (q *QueryString) String() string {
	switch v := q.value.(type) {
	case int:
		return q.key + "=" + strconv.Itoa(v)
	case string:
		return q.key + "=" + v
	default:
		return ""
	}
}

func TestBooksRoute(t *testing.T) {
	r := gin.Default()
	handlers := handler.NewApiHandler()
	router.Routes(r, handlers)

	w := httptest.NewRecorder()
	uri := &URI{uri: "/books"}
	uri = uri.withToken("hogehoge")
	// uri.withToken().withQueryString(QueryString{key: "name", value: "huga"}).
	req, _ := http.NewRequest("GET", uri.toString(), nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
