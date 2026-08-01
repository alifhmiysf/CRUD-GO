package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/models"
	"todo-api/repository"
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
		todos, err := repository.GetAllTodos()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
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
	todo, err := repository.CreateTodo(request.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
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

	todo, err := repository.GetTodoByID(id)
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

	todo, err := repository.UpdateTodo(
		id,
		request.Title,
		request.Completed,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Todo tidak ditemukan", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)

}

func DeleteTodoByID(w http.ResponseWriter, r*http.Request){
	parts := strings.Split(r.URL.Path,"/")	
	if len (parts) <3 {
		http.Error(w, "ID Tidak ditemukan", http.StatusBadRequest)
		return
	}

	id,err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "ID harus berupa angka", http.StatusBadRequest)
		return
	}

	
	err = repository.DeleteTodo(id)
	if err != nil {
		if err == sql.ErrNoRows{
			http.Error(w, "Todo tidak ditemukan", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}