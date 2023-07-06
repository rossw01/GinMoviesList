package middleware

import (
	"fmt"
	"ginTest/initialize"
	"ginTest/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
)

func CheckUserIsAuthor(c *gin.Context) {
	//idErr := c.Bind(&controllers.ReviewRequestBody)
	//if idErr != nil {
	//	return
	//}

	var targetReview models.Review
	initialize.DB.First(&targetReview, c.Param("id"))
	if targetReview.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user models.User
		initialize.DB.First(&user, claims["sub"])
		fmt.Println(user.ID, targetReview.AuthorId)
		if user.ID != uint(targetReview.AuthorId) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
