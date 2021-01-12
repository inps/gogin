package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	handledb()

	fmt.Println("web server starting  ...")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()

	fmt.Println("web server stoped!")
}
