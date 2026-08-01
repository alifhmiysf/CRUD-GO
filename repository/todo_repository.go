package repository

import (
	"database/sql"
	"todo-api/database"
	"todo-api/models"
)

func GetAllTodos() ([]models.Todo, error){
	rows,err := database.DB.Query(
		`SELECT id, title, completed FROM todos ORDER BY id`,
	)
	if err != nil {
			return nil, err 
	}
	defer rows.Close()
	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
		)
		if err != nil {
			return  nil,err
		}
		todos= append(todos, todo)
	}
	return todos,nil
}


func GetTodoByID(id int) (models.Todo, error){
	var todo models.Todo 

	err := database.DB.QueryRow (
		`SELECT id, title, completed FROM todos WHERE id = $1 `,
		id,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,	
	)

	
	if err != nil {
		return models.Todo{},err
		}
	return todo, nil
}


func CreateTodo(title string) (models.Todo, error){
	var todo models.Todo
	err := database.DB.QueryRow(
		`INSERT INTO todos(title, completed)
		VALUES($1, $2)
		RETURNING id`,
		title,
		false,
	).Scan(&todo.ID)
	
	if err != nil {
		return models.Todo{},err
	}

	todo.Title = title
	todo.Completed = false

	return  todo, nil
}


func UpdateTodo(id int, title string, completed bool)(models.Todo, error){
	
	result, err := database.DB.Exec(
		`UPDATE todos SET title = $1,completed= $2 WHERE id= $3`, 
		title,
		completed,
		id,
	)
	if err != nil {
		return models.Todo{},err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return models.Todo{}, err
    }

	if rows == 0 {
		return  models.Todo{}, sql.ErrNoRows
	}
	return models.Todo{
		ID:id,
		Title: title,
		Completed: completed,
	} ,nil	

}


func DeleteTodo(id int) error {
	result, err := database.DB.Exec(
		`DELETE FROM todos WHERE id = $1`,
		id,
	)
	if err != nil {
		return  err
	}
	rows, err :=result.RowsAffected()
	if err != nil {
		return  err
	}
	if rows == 0 {
		return  sql.ErrNoRows 
	}

	return nil
}