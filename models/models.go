package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    uint8  `json:"priority"`
	CreatedAt   int    `json:"created_at"`
	Completed   bool   `json:"completed"`
}
