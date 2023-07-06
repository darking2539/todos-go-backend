package services

import (
	"todos-go-backend/models"
	repo "todos-go-backend/repositories"
	uuid "github.com/satori/go.uuid"
)

func GetListTodosSv(userModel models.User) (*models.GetListResponse, error) {

	todos, err := repo.ListTodosData(userModel.Username)
	if err != nil {
		return nil, err
	}

	resp := models.GetListResponse{
		Data: todos,
	}
	
	return &resp, nil
}

func SubmitStatusSv(userModel models.User, request models.SubmitTodosRequest) (*models.GeneralSucessResp, error) {

	todoModel := models.Todo{
		Id: request.Id,
		Title: request.Title,
		Complete: request.Completed,
		Username: userModel.Username,
	}
	
	if todoModel.Id == "" {
		todoModel.Id = uuid.NewV4().String()
		_, err := repo.InsertTodoData(todoModel)
		if err != nil {
			return nil, err
		}
	}else {
		_, err := repo.UpdatedTodoData(todoModel)
		if err != nil {
			return nil, err
		}
	}

	resp := models.GeneralSucessResp{
		Status: "Sucessful",
		StatusCode: 200,
		Detail: todoModel.Id,
	}

	return &resp, nil
}

func DeleteTodosSv(userModel models.User, id string) (*models.GeneralSucessResp, error) {

	_, err := repo.DeleteTodoData(id, userModel.Username)
	if err != nil {
		return nil, err
	}

	resp := models.GeneralSucessResp{
		Status: "Sucessful",
		StatusCode: 200,
		Detail: id,
	}

	return &resp, nil
}