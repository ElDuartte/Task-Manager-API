package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Received /ping request")
		fmt.Fprintln(w, "pong")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Received /hello request")
		fmt.Fprintln(w, "Hello, Juan!")
	})

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
