package main

import (
	"fmt"
	"net/http"
)

type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

func main(){
	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Received /task request")

		task := Task{
			ID: 1,
			Title: "Learn Go Basics",
			Completed: false,
		}

		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(task)
	})
	
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
