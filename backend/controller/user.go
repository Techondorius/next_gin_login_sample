package controller

import (
	"backend/model"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type LoginRequest struct {
	UserID   string
	Password string
}

func Login(c *gin.Context) {
	req := LoginRequest{}
	c.BindJSON(&req)
	fmt.Println(req)

	if req.UserID == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "UserIDまたはPasswordのどちらか、または両方が空です"})
		return
	}

	user := model.FindUserByUserID(req.UserID)
	if user.Password != req.Password {
		println(user.Password, req.UserID)
		c.JSON(401, gin.H{"error": "UserIDまたはパスワードが間違っています"})
		return
	}

	token := GenerateJWT(user.UserID)
	c.SetSameSite(http.SameSiteStrictMode)
	c.SetCookie("token", token, 3600*24*7, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "OK"})
}

func GenerateJWT(uid string) string {
	claims := jwt.MapClaims{
		"userid": uid,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("asdfasdf"))
	log.Println(tokenString)
	return tokenString
}
