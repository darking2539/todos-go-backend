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

func TestSubmitTodosSv(t *testing.T) {

	request := models.SubmitTodosRequest{
		Id: "",
		Title: "ทำการบ้าน",
		Completed: false,
	}

	userModel := models.User{
		Id: "",
		Username: "",
		Email: "",
	}

	resp, err := services.SubmitStatusSv(userModel, request)
	if err != nil {
		t.Errorf("SubmitStatusSv failed with: %s", err.Error())
	}

	t.Log(resp)
}

func TestListTodosSv(t *testing.T) {

	userModel := models.User{
		Id: "",
		Username: "",
		Email: "",
	}

	resp, err := services.GetListTodosSv(userModel)
	if err != nil {
		t.Errorf("GetListTodosSv failed with: %s", err.Error())
	}

	t.Log(resp)
}

func TestDeleteSv(t *testing.T) {

	userModel := models.User{
		Id: "",
		Username: "",
		Email: "",
	}

	resp, err := services.DeleteTodosSv(userModel, "id")
	if err != nil {
		t.Errorf("DeleteTodosSv failed with: %s", err.Error())
	}

	t.Log(resp)
}

