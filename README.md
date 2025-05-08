# Task-Manager-API

RESTful API built with Golang

## Requirements:

```
go version 1.24.2
```

## To start the server:

```
go run .
```

## Endpoints URL:

### GET:

```
http://localhost:8080/tasks
```

### DELETE:

```
http://localhost:8080/task/delete?id=XXX
```
_change the XXX to the desire ID_

### POST:

```
http://localhost:8080/tasks
```

- Body JSON raw:

```
{
  "title": "Test test test"
}
```

**The API will assign an ID and completed (for now)**

## Response Example:

```
{
  "id": 1,
  "title": "Learn Docker test",
  "completed": false
},
```
