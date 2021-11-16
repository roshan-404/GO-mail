package main

import (
	"log"
	"os"

	"go-mail/config"
	"go-mail/model"
	"go-mail/routes"

	"github.com/joho/godotenv"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/go-sql-driver/mysql"
)

// @title Email services API Documentation.
// @version 1.0.0
// @description A service where users can register and send an email & do live chat.
// @termsOfService http://swagger.io/terms/

// @contact.name Roshan Kumar Ojha
// @contact.email roshankumarojha04@gmail.com

// @host localhost:3000
// @BasePath /api/v1

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&model.User{})

	r := routes.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + os.Getenv("LOCAL_PORT"))
}