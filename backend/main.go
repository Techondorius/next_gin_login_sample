package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(
		cors.Config{
			AllowAllOrigins: true,
		},
	))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello0"})
	})

	r.Run()
}
