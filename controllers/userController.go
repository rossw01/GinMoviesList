package controllers

import (
	"ginTest/initialize"
	"ginTest/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

var userRequestBody struct {
	Username string
	Password string
}

func CreateUser(c *gin.Context) {
	err := c.Bind(&userRequestBody)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userRequestBody.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Username: userRequestBody.Username, Password: string(hash)}

	result := initialize.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User '" + user.Username + "' was created successfully"})
}

func LoginUser(c *gin.Context) {
	if c.Bind(&userRequestBody) != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}
	log.Println(userRequestBody)

	var userToLogin models.User
	initialize.DB.First(&userToLogin, "username = ?", userRequestBody.Username)
	if userToLogin.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"response": "Failed to login"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(userToLogin.Password), []byte(userRequestBody.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "Failed to login",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userToLogin.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token"})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"token": tokenString})
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt",
		tokenString,
		3600*24*30, /* This is in seconds */
		"",
		"",
		false, /* set true when hosting */
		true)
	c.Status(200)
}

func GetUserByID(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	initialize.DB.First(&user, userId)

	if user.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User
	result := initialize.DB.First(&user, userId).Delete(&user)
	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"message": "User '" + user.Username + "' was deleted successfully.",
	})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	err := c.Bind(&userRequestBody)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
	}

	var user models.User
	result := initialize.DB.First(&user, userId).Updates(models.User{
		Username: userRequestBody.Username, Password: userRequestBody.Password})

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "I'm in"})
}
