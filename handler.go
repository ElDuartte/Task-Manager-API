package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Recieved /ping request")
	fmt.Fprint(w, "pong")
}

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Recieved /hello request")
	fmt.Fprint(w, "Hello!")
}

func taskHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Recieved /task request")

	if len(tasks) == 0 {
		http.Error(w, "No tasks aviable", http.StatusNotFound)
		return
	}

	firstTask := tasks[0]

	// Tell the browser we send a JSON
	w.Header().Set("Content-type", "application/json")
	// Convert struct to JSON
	json.NewEncoder(w).Encode(firstTask)
}

func tasksHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		// handle GET: Returns all tasks
		fmt.Println("Received GET /tasks request")
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	} else if r.Method == http.MethodPost {
		// handle POST: Create new task
		fmt.Println("Recieved POST /tasks request")

		// in this case Task is the type
		//----------vvvv
		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		fmt.Println("Received new task:", newTask)

		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Set ID manualy increase by 1
		newTask.ID = len(tasks) + 1
		newTask.Completed = false

		tasks = append(tasks, newTask)

		err = SaveTasks() // <--- save to file
		if err != nil {
			http.Error(w, "Failed to save task", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTask)
	} else {
		// Error method not allowed
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
