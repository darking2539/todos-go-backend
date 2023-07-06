package db

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/darahayes/go-boom"
	"github.com/gin-gonic/gin"
)

func GetConnection() (*sql.DB, error) {
	sqlConn := os.Getenv("MYSQL_URL")
	db, connErr := sql.Open("mysql", sqlConn)
	if connErr != nil {
		return nil, connErr
	}
	return db, nil
}

func Healthz(c *gin.Context) {

	err := CheckConnection()

	if err != nil {
		boom.Internal(c.Writer, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "Connect Sucessful"})
}

func CheckConnection() error {
	
	db, dbConErr := GetConnection()
	if dbConErr != nil {
		return dbConErr
	}
	defer db.Close()

	query, queryErr := db.Query("select 1")
	if queryErr != nil {
		return queryErr
	}
	defer query.Close()

	return nil
}

func InitTableDB() error {
	
	db, dbConErr := GetConnection()
	if dbConErr != nil {
		return dbConErr
	}
	defer db.Close()

	sqlUserTable := `CREATE TABLE IF NOT EXISTS users (
		id       VARCHAR(255) PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		email    VARCHAR(255) NOT NULL,
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	sqlTodosTable := `CREATE TABLE IF NOT EXISTS todos (
		id       VARCHAR(255) PRIMARY KEY,
		title    VARCHAR(255) NOT NULL,
		complete BOOL         NOT NULL,
		username VARCHAR(255) NOT NULL,
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`

	_, err := db.Exec(sqlUserTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(sqlTodosTable)
	if err != nil {
		return err
	}

	return nil
}