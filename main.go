package main

import (
	"ginTest/controllers"
	"ginTest/initialize"
	"ginTest/middleware"
	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	initialize.LoadDotEnv()
	initialize.ConnectToDB()
	middleware.CORSMiddleware(r)
}

func main() {

	// User routes
	r.POST("/register", controllers.CreateUser)
	r.POST("/login", controllers.LoginUser)

	// TODO: PREVENT TOTAL DESTRUCTION
	//r.GET("/user/:id", controllers.GetUserByID)
	//r.DELETE("/user/:id", controllers.DeleteUser)
	//r.PUT("/user/:id", middleware.CheckUserIsAuthor, controllers.UpdateUser)

	// Review routes
	r.POST("/review", controllers.CreateReview)
	r.GET("/review/:id", controllers.GetReviewsByMovieID)
	r.GET("review/user/:id", controllers.GetReviewsByAuthor)
	r.GET("/review/:id/rating", controllers.GetMovieAverageRating)
	r.DELETE("/review/:id", controllers.DeleteReview)
	r.PUT("/review/:id", middleware.CheckUserIsAuthor, middleware.RequireAuth, controllers.UpdateReview)

	r.Run()
}
