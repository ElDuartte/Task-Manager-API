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
	
	task := Task{
		ID: 1,
		Title: "Learn Go basics",
		Completed: false,
	}
	
	// Tell the browser we send a JSON
	w.Header().Set("Content-type", "application/json")
	// Convert struct to JSON
	json.NewEncoder(w).Encode(task)
}

func tasksHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received /task request")


	// Tell the browser we send a JSON
	w.Header().Set("Content-type", "application/json")
	// Convert struct to JSON
	json.NewEncoder(w).Encode(tasks)
}
