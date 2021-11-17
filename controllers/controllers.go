package controllers

import (
	queue "go-mail/hermes-mail"
	"go-mail/model"
	repo "go-mail/repository"
	"go-mail/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonlvhit/gocron"
)

type Response model.Response

// LogIn controller
// @Summary LogIn with credentials.
// @Description A registered user can login with their credentials.
// @Tags LogIn
// @Accept  json
// @Produce  json
// @Param user body model.User true "LogIn User"
// @Success 200 {object} model.User
// @Failure 401 {object} object
// @Router /login [post]
func LogIn(ctx *gin.Context) {
	var user model.User

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res := repo.SignIn(&user)
	if !res.Success {
		ctx.JSON(http.StatusUnauthorized, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// Sign Up controller
// @Summary Sign Up with credentials.
// @Description A new user can sign up with their email & password.
// @Tags Sign Up
// @Accept  json
// @Produce  json
// @Param user body model.User true "Sign Up User"
// @Success 200 {object} model.User
// @Failure 401 {object} object
// @Router /signup [post]
func SignUp(ctx *gin.Context) {
	var user model.User 

	if credErr := ctx.ShouldBindJSON(&user); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	res := repo.SignUp(&user)
	if !res.Success  {
		ctx.JSON(http.StatusNotFound, res)
		return
	}

	ctx.JSON(http.StatusOK, res)
}


// Email controller
// @Summary Varify token & send an email.
// @Description You need to signedIn and give a Token in headers then "Send Email" will execute.
// @Tags Email Compose
// @Accept  json
// @Produce  json
// @Param template body model.EmailTemplate true "Send an email"
// @Success 200 {object} model.EmailTemplate
// @Failure 401 {object} object
// @Router /compose [post]
func EmailComposer(ctx *gin.Context) {
	var T model.EmailTemplate
	
	if credErr := ctx.ShouldBindJSON(&T); credErr != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid input provided")
		return
	}

	queue.CreateQueue()

	s := gocron.NewScheduler()
	s.Every(1).Day().At("09:38").Do(queue.Dispatch, T)
	s.Start()
	log.Println("Email in progress....")
	ctx.JSON(http.StatusOK,Response{Message: "Email in progress", Data: nil, Success: true})
}


// Refresh token controller
// @Summary Varify token & create a new token.
// @Description You need to signedIn and give a Token in headers then "Refresh Token" will execute.
// @Tags Refresh token
// @Accept  json
// @Produce  json
// @Router /refreshToken [post]
func RefreshToken(ctx *gin.Context) {
	td, err := utils.VerifyRefreshToken(ctx)

	if err != "" {
		ctx.JSON(http.StatusNotFound, Response{Message: err, Data: nil, Success: false})
		return
	}
	
	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	ctx.JSON(http.StatusOK, Response{Message: "Successfully refresh token.", Data: tokens, Success: true})
}
