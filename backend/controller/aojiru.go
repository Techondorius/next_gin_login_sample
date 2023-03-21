package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Aojiru(c *gin.Context) {
	_, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(200, gin.H{"message": "Congratulations!!"})
}

func CheckCookie(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{"token": token})
}
