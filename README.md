# Todo API with Go & PostgreSQL

A simple RESTful Todo API built with **Golang** and **PostgreSQL** using the **Repository Pattern**. This project was created to learn backend development fundamentals, REST API design, database integration, and clean project structure.

## ✨ Features

- Get all todos
- Get todo by ID
- Create a new todo
- Update an existing todo
- Delete a todo
- PostgreSQL integration
- Repository Pattern
- JSON REST API
- Proper HTTP status codes
- Error handling

---

## 🛠️ Tech Stack

- Go
- PostgreSQL
- database/sql
- lib/pq

---

## 📁 Project Structure

```
todo-api/
├── database/
│   └── postgres.go
├── handlers/
│   └── todo_handler.go
├── models/
│   └── todo.go
├── repository/
│   └── todo_repository.go
├── main.go
├── go.mod
└── README.md
```

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/todo-api.git
cd todo-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create PostgreSQL database

```sql
CREATE DATABASE todo_api;
```

### 4. Create table

```sql
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);
```

### 5. Configure database connection

Edit your database configuration inside `database/postgres.go`.

Example:

```go
connStr := "host=localhost port=5432 user=postgres password=your_password dbname=todo_api sslmode=disable"
```

### 6. Run the application

```bash
go run .
```

Server will start at

```
http://localhost:8080
```

---

# API Endpoints

## Get all todos

```
GET /todos
```

Response

```json
[
  {
    "id": 1,
    "title": "Learn Golang",
    "completed": false
  }
]
```

---

## Get todo by ID

```
GET /todos/{id}
```

---

## Create todo

```
POST /todos
```

Request

```json
{
  "title": "Learn PostgreSQL"
}
```

---

## Update todo

```
PUT /todos/{id}
```

Request

```json
{
  "title": "Learn Repository Pattern",
  "completed": true
}
```

---

## Delete todo

```
DELETE /todos/{id}
```

Returns

```
204 No Content
```

---

## HTTP Status Codes

| Status | Description |
|---------|-------------|
| 200 | OK |
| 201 | Created |
| 204 | No Content |
| 400 | Bad Request |
| 404 | Not Found |
| 500 | Internal Server Error |

---

## 📚 What I Learned

- Building REST APIs using Go
- Working with `database/sql`
- CRUD operations with PostgreSQL
- Repository Pattern
- HTTP request and response handling
- JSON encoding and decoding
- Error handling
- Clean project structure

---

## 🔮 Future Improvements

- Service Layer
- JWT Authentication
- Middleware
- Environment Variables (.env)
- Docker
- Unit Testing
- API Documentation (Swagger)
- Deploy to Railway or VPS

---

## 👨‍💻 Author

Ali Fahmi Yusuf

GitHub: https://github.com/alifhmiysf