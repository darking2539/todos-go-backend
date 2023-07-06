package repositories

import (
	"errors"
	"fmt"
	"todos-go-backend/db"
	"todos-go-backend/models"
)

func CountUserData(username string, email string) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	var queryNum int64 = 0
	sqlStr := fmt.Sprintf("select count(0) from users where username='%s' or email='%s' ", username, email)

	stmt, stmtErr := conn.Prepare(sqlStr)
	if stmtErr != nil {
		return nil, stmtErr
	}
	
	queryErr := stmt.QueryRow().Scan(&queryNum)
	if queryErr != nil {
		return nil, queryErr
	}

	return &queryNum, nil
}

func InsertUserData(uuid string, req models.RegisterRequest) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("insert into users(id, username, email, password) values('%s', '%s', '%s', '%s')", uuid, req.Username, req.Email, req.Password)

	stmt, prepErr := conn.Prepare(sqlStr)
	if prepErr != nil {
		return nil, prepErr
	}

	resp, execErr := stmt.Exec()
	if execErr != nil {
		return nil, execErr
	}

	numRows, numRowsErr := resp.RowsAffected()
	if numRowsErr != nil {
		return nil, numRowsErr
	}
	
	
	return &numRows, nil
}

func GetUserData(username string) (*models.User, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("select id, username, password, email from users where username='%s'", username)

	stmt, stmtErr := conn.Prepare(sqlStr)
	if stmtErr != nil {
		return nil, stmtErr
	}
	
	var item models.User
	scanErr := stmt.QueryRow().Scan(&item.Id, &item.Username, &item.Password, &item.Email)
	if scanErr != nil {
		return nil, scanErr
	}

	return &item, nil
}

func UpdatedPassword(id string, passwordHash string) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("UPDATE users SET password='%s' where id='%s' ", passwordHash, id)
	fmt.Println(sqlStr)

	stmt, prepErr := conn.Prepare(sqlStr)
	if prepErr != nil {
		return nil, prepErr
	}

	resp, execErr := stmt.Exec()
	if execErr != nil {
		return nil, execErr
	}

	numRows, numRowsErr := resp.RowsAffected()
	if numRowsErr != nil {
		return nil, numRowsErr
	}

	if numRows == 0 {
		err := errors.New("password not change")
		return nil, err
	}
	
	return &numRows, nil
}