package main

import "time"

type Task struct {
    ID           int        `json:"id"`
    CreatorID    int        `json:"creatorId"`
		ProjectID    []int      `json:"projectId,omitempty"`
		AssignedToID []int      `json:"assignedToId,omitempty"`
    Title        string     `json:"title"`
    Status       string     `json:"status"`
    Deleted      bool       `json:"deleted"`
    Edited       time.Time  `json:"time"`
		CreatedTime  time.Time  `json:"createdTime"`
}

// ProjectID []int is the syntax to declare an array of integers
// AssignedToID

