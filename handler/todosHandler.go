package handler

import (
	"encoding/json"
	"todos-go-backend/models"
	sv "todos-go-backend/services"
	"todos-go-backend/utils"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func GetListTodosHandler(c *gin.Context) {

	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}

	resp, svErr := sv.GetListTodosSv(*userModel)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func SubmitTodosHandler(c *gin.Context) {
	
	var request models.SubmitTodosRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}

	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}
	
	resp, svErr := sv.SubmitStatusSv(*userModel, request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func DeleteTodosHandler(c *gin.Context) {

	taskId := c.Param("id")
	if taskId == "" {
		boom.BadRequest(c.Writer, "id is Empty");
		return
	}

	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}
	
	resp, svErr := sv.DeleteTodosSv(*userModel, taskId)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}