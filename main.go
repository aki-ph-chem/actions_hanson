package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, makeGreeting(name))
	})
	router.Run(":8080")
}

func makeGreeting(name string) string {
	return fmt.Sprintf("Hello,%s", name)
}
