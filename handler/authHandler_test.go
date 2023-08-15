package handler

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"todos-go-backend/db"
	"todos-go-backend/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	godotenv.Load("../.env")

	//health db is connect
	err := db.CheckConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	//initz table db
	err = db.InitTableDB()
	if err != nil {
		log.Println(err)
	}
}

func TestRegisterHandler(t *testing.T) {
	gin.SetMode(gin.TestMode) // Set Gin to test mode

	// Create a new router and register the route
	router := gin.Default()
	router.POST("/register", RegisterHandler)

	// Test case 1: Valid request
	validPayload := models.RegisterRequest{
		Username: "abosszzzzz",
		Password: "123456222",
		Email:    "suphadetttt.b@gmail.com",
	}

	validPayloadJSON, _ := json.Marshal(validPayload)
	validReq, err := http.NewRequest("POST", "/register", bytes.NewBuffer(validPayloadJSON))
	if err != nil {
		t.Error(err)
	}
	validResp := httptest.NewRecorder()
	router.ServeHTTP(validResp, validReq)
	if !assert.Equal(t, http.StatusOK, validResp.Code) {
		t.Errorf("Body is: %s", validResp.Body)
	}

	// Test case 2: Empty field in the request
	emptyFieldPayload := models.RegisterRequest{
		// Fill in payload with empty field here
	}
	emptyFieldPayloadJSON, _ := json.Marshal(emptyFieldPayload)
	emptyFieldReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(emptyFieldPayloadJSON))
	emptyFieldResp := httptest.NewRecorder()
	router.ServeHTTP(emptyFieldResp, emptyFieldReq)
	assert.Equal(t, http.StatusBadRequest, emptyFieldResp.Code)

}
