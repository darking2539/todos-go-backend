package handler

import (
	"encoding/json"
	"todos-go-backend/models"
	sv "todos-go-backend/services"
	"todos-go-backend/utils"

	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func RegisterCtl(c *gin.Context) {
	
	var request models.RegisterRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, "field can't empty");
		return
	}
	
	resp, svErr := sv.RegisterSv(request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func LoginCtl(c *gin.Context) {

	var request models.LoginRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, "field can't empty");
		return
	}
	
	resp, svErr := sv.LoginSv(request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}

func ChangePasswordCtl(c *gin.Context) {
	
	var request models.ChangePasswordRequest
	if payloadErr := c.ShouldBindJSON(&request); payloadErr != nil {
		boom.BadRequest(c.Writer, payloadErr.Error());
		return
	}

	userModel, err := utils.GetUserInfo(c)
	if err != nil {
		boom.BadRequest(c.Writer, err.Error());
		return
	}
	
	resp, svErr := sv.ChangePasswordSv(*userModel, request)
	if svErr != nil {
		boom.BadRequest(c.Writer, svErr.Error());
		return
	}

	c.Writer.Header().Set("Content-Type", "application/json");
	json.NewEncoder(c.Writer).Encode(&resp);
}