package main 
import (
	"fmt"
	"net/http"
	"strings"
	"todo-api/handlers"
	"todo-api/database"
) 
func main(){
	database.Connect()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		
			if r.URL.Path == "/todos" {
					switch r.Method {
				case http.MethodGet:
					handlers.GetTodos(w,r)

				case http.MethodPost:
					handlers.CreateTodo(w,r)

				default:
					http.Error(w,"Method tidak diizinkan", http.StatusMethodNotAllowed)
				}
				return
			}
		
		
			if strings.HasPrefix(r.URL.Path, "/todos/"){
				switch r.Method {

				case http.MethodGet:
				handlers.GetTodoByID(w, r)
				
				case http.MethodPut:
					handlers.UpdateTodo(w,r)

				case http.MethodDelete:
					handlers.DeleteTodoByID(w,r)
					
				default:
				http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
					}
				return
			}
		http.NotFound(w,r)
	})
	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}