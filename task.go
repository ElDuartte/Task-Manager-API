package main
import "time"

type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Completed bool      `json:"completed"`
    Deleted   bool      `json:"deleted"`
    Edited    time.Time    `json:"time"`
    CreatedTime time.Time 
    AssignedToID int
    
}

# Creatorid
# ProjectID -- needs to be an array in case it belongs to multiple projects at the same time
# AssignedToID -- needs to be an array in case it's for multiple devs

