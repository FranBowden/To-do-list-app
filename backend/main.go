//Task:
/*Create the foundation of a to-do list application, focusing on backend functionality and essential frontend interaction.
Your task is implementing a RESTful API using Go and a simple JS interface using React.
The goal is to be able to be able to list and add todo's.
The backend needs to meet the openapi spec which within the backend folder. You need to create the list endpoint and the add todo endpoint.
The storage system is in-memory.
The frontend already has functionality to list the todo's, your task is to complete the form which submits todo's to the backend system.*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Struct for Todo list that contains a title and description, represents a to do item
type ToDo struct {
	Title string  `json:"title"`
	Description string	`json:"description"`
}

//Store the to dos
//var  ToDos []ToDo //hold the list of todos

// This is an example of a todo that list the "todo"
// unsure about the openapi spec and how that gets implemented
 var ToDos = []ToDo{
        {Title: "Go for a jog", Description: "run 4km around the river"}, //example
		  {Title: "test", Description: "this is a test example"},
    }


func main() {
	// Your code here
	
	//start server on port 8080
	http.HandleFunc("/", ToDoListHandler) 
	http.ListenAndServe(":8080", nil)
	

}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// Your code here
	switch r.Method {
		case "GET":
			fmt.Print("GET")
			listToDos(w, r) //get request
		case "POST":
			fmt.Print("POST")
			addTodo(w, r) //post request
		default:
			http.Error(w, "invalid method", http.StatusMethodNotAllowed) //error message
	}

}


func addTodo(w http.ResponseWriter, r *http.Request) {

	var newToDo ToDo

	if err := json.NewDecoder(r.Body).Decode(&newToDo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Adding todo:", newToDo)

	ToDos = append(ToDos, newToDo) //add the to do to the list

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newToDo)
}


func listToDos(w http.ResponseWriter,  r *http.Request) {
//	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ToDos) //returns the list of todos as json

	log.Printf("Todos: %+v", ToDos) //debugging
}