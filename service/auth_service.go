package service

import (
	"fmt"
	"go-jwt/db"
	"go-jwt/jwt"
	"go-jwt/models"
	"go-jwt/params"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CurrentUser(c *gin.Context) *models.User {
	sub, ok := c.Get("sub")
	if !ok {
		return &models.User{}
	}
	return db.FindUserById(fmt.Sprintf("%s", sub))
}

func AuthPingHandler(c *gin.Context) {
	user := CurrentUser(c)
	c.String(http.StatusOK, fmt.Sprintf("auth pong: %s", user.Email))
}

func ChangePasswordHandler(c *gin.Context) {
	sub, ok := c.Get("sub")
	if !ok {
		return
	}

	var change params.ChangePassword
	if err := c.ShouldBind(&change); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if change.Password != change.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "password and password confirm not match"})
		return
	}

	user := db.FindUserById(fmt.Sprintf("%v", sub))
	err := bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(change.OriginPassword))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "origin password error"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(change.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedPassword := string(hash)
	db.DB.Model(&user).Updates(models.User{EncryptedPassword: encryptedPassword})
	payload := jwt.GenPayload("user", user.ID.String())
	jwt.RevokeLastJwt(payload)
	c.JSON(http.StatusOK, gin.H{"status": "update password success"})
}
