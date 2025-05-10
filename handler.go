package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

		// Set ID, Status, Description
		newTask.ID = len(tasks) + 1
		newTask.Status = "To Do"
		newTask.Description = ""

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

func deleteTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Recieved DELETE /tasks/delete request")

	// Parse task ID from query: /task/delete?id=XXX
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	// Atoi is "ASCII to Integer"
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	// find task and mark it ask deleted
	for i := range tasks {
		if tasks[i].ID == id{
			tasks[i].Deleted = true
			tasks[i].Edited = time.Now()

			err = SaveTasks()
			if err != nil{
				http.Error(w, "Failed to save deletion", http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Task %d soft-deleted", id)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func editTaskHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPut{
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path: /tasks?id=3
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// find the task
	// *Task means is a pointer to a Task, not the struct itself
	// Pointers let you edit the original data and not a copy
	var taskToUpdate *Task
	for i := range tasks {
		if tasks[i].ID == id{
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil{
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// decode the data from the request
	var updatedData Task
	// &updatedData the syntax "&" is indicating it's a pointer to the variable
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// update fields
	// Its uggly but it's honest work...
	if updatedData.Title != ""{
		taskToUpdate.Title = updatedData.Title
	}

	if updatedData.Description != "" {
		taskToUpdate.Description = updatedData.Description
	}

	if updatedData.Status != "" {
		taskToUpdate.Status = updatedData.Status
	}

	if !updatedData.Deadline.IsZero() {
		taskToUpdate.Deadline = updatedData.Deadline
	}

	if len(updatedData.AssignedToID) > 0 {
		taskToUpdate.AssignedToID = updatedData.AssignedToID
	}

	if len(updatedData.ProjectID) > 0 {
		taskToUpdate.ProjectID = updatedData.ProjectID
	}

	if updatedData.CreatorID != 0 {
		taskToUpdate.CreatorID = updatedData.CreatorID
	}

	taskToUpdate.Status = updatedData.Status
	taskToUpdate.Edited = time.Now()

	// save changes to file
	err = SaveTasks()
	if err != nil {
		http.Error(w, "Failed to save task", http.StatusInternalServerError)
		return
	}

	// RRespond with updated task
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(taskToUpdate)
}
