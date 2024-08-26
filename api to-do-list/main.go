package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

var todos []Todo

type Todo struct{
	ID		int		`json:"id"`
	Task	string	`json:"task"`
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/todos", GetTodos).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id:[0-9]+}", DeleteTodo).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request){
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo.ID = len(todos) +1
	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for index, todo := range todos {
		if todo.ID ==id{
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}
