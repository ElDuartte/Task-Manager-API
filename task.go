package main

import "time"

type Task struct {
    ID        int          `json:"id"`
    CreatorID int          `json:"creatorId"`
		ProjectID []int        `json:"projectId"`
		AssignedToID []int     `json:"assignedToId"`
    Title     string       `json:"title"`
    Completed bool         `json:"completed"`
    Deleted   bool         `json:"deleted"`
    Edited    time.Time    `json:"time"`
		CreatedTime time.Time  `json:"createdTime"`
}

// ProjectID []int is the syntax to declare an array of integers
// AssignedToID

