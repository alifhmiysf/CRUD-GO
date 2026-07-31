package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo-api/database"
	"todo-api/models"
)

var todos = []models.Todo{
	{
		ID: 1,
		Title: "Belajar Golang",
		Completed: false,
	},
	{
		ID: 2,
		Title: "Belajar REST API",
		Completed: true,
	},

}



func GetTodos(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(todos)

	// menggunakan db
	rows,err := database.DB.Query(
		"SELECT id, title, completed FROM todos ORDER BY id",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// function untuk menerima respon
func CreateTodo(w http.ResponseWriter, r *http.Request){
	var request models.CreateTodoRequest 
	err := json.NewDecoder(r.Body).Decode(&request)
	if err !=nil {
		http.Error(w, "Request tidak valid", http.StatusBadRequest)
		return 
	}

	// newTodo := models.Todo{
	// 	ID: len(todos)+1,
	// 	Title: request.Title,
	// 	Completed: false,
	// }
	// todos = append(todos, newTodo)

	var newTodo models.Todo
	err =database.DB.QueryRow(
		`INSERT INTO todos(title, completed)
		VALUES($1, $2)
		RETURNING id`,
		request.Title,
		false,
	).Scan(&newTodo.ID)
	fmt.Println("ID:", newTodo.ID)
	fmt.Println("ERR:", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newTodo.Title = request.Title
	newTodo.Completed = false


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTodo)
}

// function untuk mengambil per id 
func GetTodoByID(w http.ResponseWriter, r *http.Request){
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) <3 {
		http.Error(w,"ID Tidak ditemukan", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w,"ID harus berupa angka", http.StatusBadRequest)
		return
	}
	var todo models.Todo 
	err = database.DB.QueryRow (
		`SELECT id, title, completed FROM todos WHERE id = $1 `,
		id,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,	
	)
	fmt.Println("ID:", id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo tidak ditemukan", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}


func UpdateTodo(w http.ResponseWriter, r*http.Request){
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) <3 {
		http.Error(w,"ID Tidak ditemukan ", http.StatusBadRequest)
		return
	}
	var request models.UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Request tidak valid", http.StatusBadRequest)
		return
	}


	id,err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	// for i := range todos {
	// 	if todos[i].ID == id {
			
	// 		todos[i].Title = request.Title
	// 		todos[i].Completed= request.Completed
			
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(todos[i])
	// 		return
	// 	}
	// }

	result, err := database.DB.Exec(
		`UPDATE todos SET title = $1,completed= $2 WHERE id= $3`, 
		request.Title,
		request.Completed,
		id,
	)
	rows, err := result.RowsAffected()
	if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
    }
	

	if rows == 0 {
		http.Error(w, "Todo tidak ditemukan", http.StatusNotFound)
		return
	}
}

func DeleteTodoByID(w http.ResponseWriter, r*http.Request){
	parts := strings.Split(r.URL.Path,"/")	
	if len (parts) <3 {
		http.Error(w, "ID Tidak ditemukan", http.StatusBadRequest)
		return
	}
	var request models.UpdateTodoRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Request tidak valid", http.StatusBadRequest)
		return
	}


	id,err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}