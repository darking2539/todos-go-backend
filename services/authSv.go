package services

import (
	"errors"
	"todos-go-backend/middleware"
	"todos-go-backend/models"
	repo "todos-go-backend/repositories"
	"todos-go-backend/utils"

	uuid "github.com/satori/go.uuid"
)

func LoginSv(request models.LoginRequest) (*models.LoginResponse, error) {
	
	userData, err := repo.GetUserData(request.Username)
	if err != nil {
		
		if err.Error() == "sql: no rows in result set" {
			err = errors.New("username not found")
			return nil, err
		}else {
			return nil, err
		}

	}

	if !utils.CheckPasswordHash(request.Password, userData.Password) {
		err = errors.New("wrong password")
		return nil, err
	}

	jwtToken := middleware.JwtSign(*userData)

	resp := models.LoginResponse{
		Token: jwtToken,
	}
	
	return &resp, nil
}

func RegisterSv(request models.RegisterRequest) (*models.GeneralSucessResp, error) {

	uuid := uuid.NewV4().String()

	count, err := repo.CountUserData(request.Username, request.Email)
	if err != nil {
		return nil, err
	}

	if count != nil && *count != 0 {
		err = errors.New("user already exists")
		return nil, err
	}

	//encrypt password
	passwordHash, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	//replace data
	request.Password = passwordHash

	_, err = repo.InsertUserData(uuid, request)
	if err != nil {
		return nil, err
	}

	resp := models.GeneralSucessResp{
		Status: "Sucessful",
		StatusCode: 200,
		Detail: "Register Sucessful",
	}

	return &resp, nil
}

func ChangePasswordSv(usermodel models.User, request models.ChangePasswordRequest) (*models.GeneralSucessResp, error) {
	
	userData, err := repo.GetUserData(request.Username)
	if err != nil {
		
		if err.Error() == "sql: no rows in result set" {
			err = errors.New("username not found")
			return nil, err
		}else {
			return nil, err
		}

	}

	if !utils.CheckPasswordHash(request.OldPassword, userData.Password) {
		err = errors.New("wrong password")
		return nil, err
	}

	hashPassword, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return nil, err
	}

	_, err = repo.UpdatedPassword(userData.Id, hashPassword)
	if err != nil {
		return nil, err
	}

	resp := models.GeneralSucessResp{
		Status: "Sucessful",
		StatusCode: 200,
		Detail: userData.Id,
	}

	return &resp, nil
}
