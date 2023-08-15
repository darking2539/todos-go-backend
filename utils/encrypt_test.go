package utils

import (
	"testing"
)

func TestHashFunction(t *testing.T) {

	password := "123456"
	
	hashPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Failed with: %s", err.Error())
	}
	
	if !CheckPasswordHash(password, hashPassword) {
		t.Errorf("Failed with: password not match" )
	}

}