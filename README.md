# WIP

# I'm still working on this project

---

# Task-Manager-API

RESTful API built with Golang

## Requirements

```
go version 1.24.2
```

## To start the server

```
go run .
```

## Endpoints URL

### GET -- list

```
http://localhost:8080/tasks
```

### DELETE -- delete single task

```
http://localhost:8080/task/delete?id=XXX
```

_change the XXX to the desire ID_

### POST -- create

```
http://localhost:8080/tasks
```

- Body JSON raw:

```
{
  "title": "Test test test"
}
```

### PUT -- edit
```
http://localhost:8080/task/edit?id=5
```
**The API will assign an ID and completed (for now)**

## Response Example

```
{
    "id": 5,
    "creatorId": 1,
    "projectId": [
      10
    ],
    "assignedToId": [
      2,
      3
    ],
    "title": "This is the updated title",
    "description": "Description for the 5 task editing",
    "status": "In Progress",
    "deleted": false,
    "deadlineTime": "2024-12-01T17:00:00Z",
    "editedTime": "2025-05-10T23:23:01.198121712+02:00",
    "createdTime": "0001-01-01T00:00:00Z"
  }
```
