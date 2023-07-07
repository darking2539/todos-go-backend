package test

import (
	"testing"
	"todos-go-backend/utils"
)

func TestHashFunction(t *testing.T) {

	password := "123456"
	
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		t.Errorf("Failed with: %s", err.Error())
	}
	
	if !utils.CheckPasswordHash(password, hashPassword) {
		t.Errorf("Failed with: password not match" )
	}

}