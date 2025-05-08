package main

import (
	"fmt"
	"net/http"
)

func main(){
	err := LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/task/delete", deleteHandler)

	fmt.Println("Server is running on port 8080")

	servErr := http.ListenAndServe(":8080", nil)

	if servErr != nil {
		fmt.Println("Error starting server:", servErr)
	}
}
