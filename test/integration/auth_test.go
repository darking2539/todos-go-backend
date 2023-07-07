package test

import (
	"testing"
	"todos-go-backend/models"
	"todos-go-backend/services"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
}

func TestRegisterSv(t *testing.T) {

	request := models.RegisterRequest{
		Username: "admin",
		Password: "123456",
		Email: "suphadet.b@gmail.com",
	}

	resp, err := services.RegisterSv(request)
	if err != nil {
		t.Errorf("Register failed with: %s", err.Error())
	}

	t.Log(resp)
	
}
func TestLoginSv(t *testing.T) {

	request := models.LoginRequest{
		Username: "admin",
		Password: "123456",
	}

	resp, err := services.LoginSv(request)
	if err != nil {
		t.Errorf("Login failed with: %s", err.Error())
	}

	t.Log(resp)
}

func TestChangePasswordSv(t *testing.T) {

	request := models.ChangePasswordRequest{
		Username: "admin",
		OldPassword: "123456",
		NewPassword: "12345678",
	}

	userModel := models.User{
		Id: "", //need this
		Username: "", //need this
		Email: "",
	}

	resp, err := services.ChangePasswordSv(userModel, request)
	if err != nil {
		t.Errorf("Login failed with: %s", err.Error())
	}

	t.Log(resp)
}

