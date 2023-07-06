package main

import (
	"fmt"
	"log"
	"os"
	"todos-go-backend/db"
	"todos-go-backend/middleware"
	"todos-go-backend/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

	godotenv.Load()
	port := os.Getenv("PORT")

	//health db is connect
	err := db.CheckConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	//initz table db
	err = db.InitTableDB()
	if err != nil {
		fmt.Println(err)
	}

	//initz gin
	engine := gin.Default()
	engine.Use(middleware.CORSMiddleware())
	router.Setup(engine)

	engine.Run(":" + port)

}
