package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("hello")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
