package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/hello/:name",func(c *gin.Context){
        name := c.Param("name")
        c.String(http.StatusOK, makeGreeting(name))
    })
}

func makeGreeting(name string) string  {
    return fmt.Sprintf("Hello,%d", name)
}
