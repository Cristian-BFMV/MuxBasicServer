package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}

func main() {
	//Creating a new Mux Router instance
	router := mux.NewRouter().StrictSlash(true)
	//Creating the server routes
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/tasks", GetTasks)
	//Running the HTTP server on port 8080
	http.ListenAndServe(":8080", router)
	fmt.Println("Server running on port 8080")
}

// HomeHandler is a Route handler function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//HTTP status code
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world")
}

// GetTasks is a Route handler function
func GetTasks(w http.ResponseWriter, r *http.Request) {
	//Setting the Content type header to application JSON
	w.Header().Set("Content-Type", "application/json")
	//Returns the tasks as a JSON
	json.NewEncoder(w).Encode(tasks)
}

// CreateTask is a Route handler function
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	//Getting the request data with the ioutil module
	data, err := ioutil.ReadAll(r.Body)
	//Checking if there is an error
	if err != nil {
		fmt.Fprintf(w, "Insert valid data")
	}
	//Converting the data to a new task
	json.Unmarshal(data, &newTask)
	//Setting task id with the length of the slice + 1
	newTask.ID = len(tasks) + 1
	// Inserting the newTask to the tasks slice
	tasks = append(tasks, newTask)
	//Setting the response code
	w.WriteHeader(http.StatusCreated)
	//Setting the Content type header to application JSON
	w.Header().Set("Content-Type", "application/json")
	//Returning the new task
	json.NewEncoder(w).Encode(newTask)
}
