package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/task", taskHandler)
	http.HandleFunc("/tasks", tasksHandler)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
