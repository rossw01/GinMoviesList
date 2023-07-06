package controllers

import (
	"ginTest/initialize"
	"ginTest/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var reviewRequestBody struct {
	// TODO: Change capitalisation
	AuthorId uint
	MovieId  string
	Rating   float32
	Subject  string
	Body     string
}

func CreateReview(c *gin.Context) {
	err := c.Bind(&reviewRequestBody)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	var user models.User
	initialize.DB.First(&user, reviewRequestBody.AuthorId)
	if user.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	review := models.Review{
		AuthorId: int(reviewRequestBody.AuthorId),
		Author:   user.Username,
		MovieId:  reviewRequestBody.MovieId,
		Rating:   reviewRequestBody.Rating,
		Subject:  reviewRequestBody.Subject,
		Body:     reviewRequestBody.Body,
		Date:     time.Now().Format("02/01/2006 15:04:05"),
	}
	result := initialize.DB.Create(&review)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(200, gin.H{
		"message": "Review with subject '" + review.Subject + "' was created successfully"})
}

func GetReviewsByMovieID(c *gin.Context) {
	movieId := c.Param("id")
	movieIdQuery := models.Review{MovieId: movieId}

	var reviews []models.Review
	initialize.DB.Find(&reviews, movieIdQuery)

	c.JSON(200, gin.H{
		"results": reviews})
}

func GetReviewsByAuthor(c *gin.Context) {
	author := c.Param("id")
	authorQuery := models.Review{Author: author}

	var reviews []models.Review
	initialize.DB.Find(&reviews, authorQuery)

	c.JSON(200, gin.H{
		"results": reviews})
}

func DeleteReview(c *gin.Context) {
	reviewId := c.Param("id")

	var reviewToDelete models.Review
	result := initialize.DB.First(&reviewToDelete, reviewId).Delete(&reviewToDelete)
	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"message": "Removed review with ID '" + reviewId + "' successfully."})
}

func UpdateReview(c *gin.Context) {
	reviewId := c.Param("id")

	idErr := c.Bind(&reviewRequestBody)
	if idErr != nil {
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	var review models.Review
	result := initialize.DB.First(&review, reviewId).Updates(models.Review{
		Rating:  reviewRequestBody.Rating,
		Subject: reviewRequestBody.Subject,
		Body:    reviewRequestBody.Body,
		Date:    time.Now().Format("02/01/2006 15:04:05"),
	})

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"review": review,
	})
}

func GetMovieAverageRating(c *gin.Context) {
	movieId := c.Param("id")

	var reviews []models.Review
	initialize.DB.Where(&models.Review{MovieId: movieId}).Find(&reviews)
	if len(reviews) == 0 {
		c.JSON(200, gin.H{
			"rating": -1})
	}

	var ratingTotal float32
	for _, review := range reviews {
		ratingTotal += review.Rating
	}

	c.JSON(200, gin.H{
		"rating": ratingTotal / float32(len(reviews)),
	})
}
