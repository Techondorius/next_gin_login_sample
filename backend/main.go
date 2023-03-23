package main

import (
	"backend/controller"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Cors
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		MaxAge:           300,
	}
	config.AddAllowHeaders(
		"Access-Control-Allow-Credentials",
		"Access-Control-Allow-Headers",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
	)
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello0"})
	})

	r.POST("/login", controller.Login)
	r.GET("/checkCookie", controller.CheckCookie)

	auth := r.Group("auth")
	auth.Use(CheckJWT)

	auth.GET("/aojiru", controller.Aojiru)

	r.Run()
}

func CheckJWT(c *gin.Context) {
	fmt.Println(c.Request.Header["Cookie"])
	tokenCookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorised"})
		c.Abort()
		return
	}
	if tokenCookie == "" {
		c.JSON(401, gin.H{"error": "unauthorised"})
		c.Abort()
		return
	}

	c.Next()
}
