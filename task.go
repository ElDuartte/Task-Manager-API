package main

import "time"

type Task struct {
    ID           int        `json:"id"`
    CreatorID    int        `json:"creatorId"`
    ProjectID    []int      `json:"projectId,omitempty"`
    AssignedToID []int      `json:"assignedToId,omitempty"`
    Title        string     `json:"title"`
    Description  string     `json:"description"`
    Status       string     `json:"status"`
    Deleted      bool       `json:"deleted"`
    Deadline     time.Time  `json:"deadlineTime"`
    Edited       time.Time  `json:"editedTime"`
    CreatedTime  time.Time  `json:"createdTime"`
}

// ProjectID []int is the syntax to declare an array of integers
// AssignedToID

