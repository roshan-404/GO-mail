package repository

import (
	"go-mail/config"
	"go-mail/model"
	"go-mail/utils"

	"golang.org/x/crypto/bcrypt"
)

type Response model.Response

var err error

func SignIn(user *model.User) Response {
	var pswd = user.Password
	if err = config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		return Response{Message: "Record not found!", Data: nil, Success: false}
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pswd)); err != nil {
		return Response{Message: "Wrong password!", Data: nil, Success: false}
	}

	// genrating JWT token
	ts, err := utils.CreateToken(user.Email)
	if err != nil {
		return Response{Message: "Something went wrong!", Data: nil, Success: false}
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	return Response{Message: "SignedIn successfully", Data: tokens, Success: true}
}

func SignUp(user *model.User) Response {
	if err = config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		//hashing password before stroing
		hash, hashErr := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if hashErr != nil {
			return Response{Message: "Something went wrong!", Data: nil, Success: false}
		}

		user.Password = string(hash)

		if err = config.DB.Create(user).Error; err != nil {
			return Response{Message: "SignedUp failed!", Data: nil, Success: false}
		}
		return Response{Message: "SignedUp successfully", Data: nil, Success: true}
	}
	return Response{Message: "Record already exists!", Data: nil, Success: false}
}
