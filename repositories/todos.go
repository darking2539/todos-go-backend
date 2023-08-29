package repositories

import (
	"errors"
	"fmt"
	"todos-go-backend/db"
	"todos-go-backend/models"
)

func InsertTodoData(req models.Todo) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("insert into todos(id, title, complete, username) values('%s', '%s', %t, '%s')", req.Id, req.Title, req.Complete, req.Username)


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

func UpdatedTodoData(req models.Todo) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("update todos set title='%s', complete=%t where id='%s' and username='%s' ", req.Title, req.Complete, req.Id, req.Username)

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
		err := errors.New("data not change")
		return nil, err
	}
	
	return &numRows, nil
}

func DeleteTodoData(id string, username string) (*int64, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("DELETE FROM todos WHERE id='%s' and username='%s'", id, username)

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
		err := errors.New("data not found")
		return nil, err
	}
	
	return &numRows, nil
}

func ListTodosData(username string) ([]models.Todo, error) {
	
	conn, dbConErr := db.GetConnection()
	if dbConErr != nil {
		return nil, dbConErr
	}
	defer conn.Close()

	sqlStr := fmt.Sprintf("SELECT id, title, complete, username FROM todos WHERE username='%s' ORDER BY created_date", username)

	stmt, stmtErr := conn.Prepare(sqlStr)
	if stmtErr != nil {
		return nil, stmtErr
	}
	
	queryRows, rowsErr := stmt.Query()
	if rowsErr != nil {
		return nil, rowsErr
	}
	defer queryRows.Close()

	items := []models.Todo{}

	for queryRows.Next() {
		var item models.Todo
		queryRows.Scan(&item.Id, &item.Title, &item.Complete, &item.Username)
		items = append(items, item)
	}

	return items, nil
}