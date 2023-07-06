package utils

import (
	"encoding/json"
	"todos-go-backend/models"

	"github.com/gin-gonic/gin"
)


func GetUserInfo(c *gin.Context) (*models.User, error) {
	
	var userInfo models.User
	userHeader := c.Request.Header.Get("X-User")

	err := json.Unmarshal([]byte(userHeader), &userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}