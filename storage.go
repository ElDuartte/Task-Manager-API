package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var tasks []Task

func LoadTasks() error {
	file, err := os.Open("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			tasks = []Task{} // if no file, start with empty arr tasks
			return nil
		}
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return err
	}

	fmt.Println("Loaded tasks from file")
	return nil
}

func SaveTasks() error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}

	fmt.Println("Saved tasks to file")
	return nil
}
