# dartz

dartz is a task management REST API built with Golang and the Gin framework.

## Features

- User registration and authentication system using JWT tokens
- CRUD operations for tasks (Create, Read, Update, Delete)
- Secure password hashing
- MySQL database integration

## API Endpoints

| Method | Endpoint          | Description                    | Authentication Required |
|--------|------------------ |------------------------------- |-------------------------|
| POST   | /register         | Registers a new user           | No                      |
| POST   | /login            | Authenticates a user           | No                      |
| GET    | /api/tasks        | Retrieves all tasks            | Yes                     |
| POST   | /api/tasks        | Creates a new task             | Yes                     |
| GET    | /api/tasks/{id}   | Retrieves a specific task      | Yes                     |
| PUT    | /api/tasks/{id}   | Updates a specific task        | Yes                     |
| DELETE | /api/tasks/{id}   | Deletes a specific task        | Yes                     |

## Installation & Usage

1. Clone and navigate to the repository:
```bash
git clone https://github.com/MakayaYoel/dartz.git
cd dartz
```

2. Rename the ``env.example`` file to ``.env``, then fill in your DB details as well as your JWT secret key.

3. Start the server. The application will start at `http://localhost:8080`:
```bash
go run cmd/main.go
```