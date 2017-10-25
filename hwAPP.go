package hello

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func init() {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	
	http.Handle("/", r)
}