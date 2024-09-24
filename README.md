# dartz

dartz is a task management REST API made with Golang and the Gin framework.

## Features
- **Authentication System:** Has an authentication system that utilizes JWT authentication tokens.
- **Configurable:** You can easily input your database details and JWT secret key.

## API Endpoints
| Methods  | Endpoints | Description |
| ------------- | ------------- | ------------- |
| GET  | /api/tasks  | Returns all tasks  |
| POST | /api/tasks  | Creates a task |
| GET  | /api/tasks/{id} | Returns all information on the specified task |
| PUT  | /api/tasks/{id} | Updates the specified task |
| DELETE | /api/tasks/{id} | Deletes the specified task |
| POST | /register | Creates a new user |
| POST | /login | Authenticates a user |
  
## Installation & Usage
- Coming Soon