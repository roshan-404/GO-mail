package routes

import (
	"go-mail/controllers"

	"os"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)


func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	apiV1 := router.Group("/" + apiPrefix + "/" + apiVersion)

	// Apply the middleware to the router (works with groups too)
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))
	
	auth := apiV1.Group("/")
	{
		auth.POST("/login", controllers.LogIn)
		auth.POST("/signup",controllers.SignUp)
	}

	secure := apiV1.Group("/")
	{
		secure.POST("/refreshToken", controllers.RefreshToken)
		secure.POST("/compose", controllers.EmailComposer)
	}
	
	

	return router
}
