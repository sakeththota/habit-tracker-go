package types

import (
	"time"
)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(user User) error
}

type HabitStore interface {
	GetHabits(user_id int) ([]Habit, error)
	GetHabitById(habit_id int) (*Habit, error)
	CreateHabit(habit Habit) error
	DeleteHabit(user_id int, habit_id int) error
}

type ProgressStore interface {
	GetProgress(habit_id int) ([]ProgressEntry, error)
	MarkComplete(habit_id int) error
}

type ProgressEntry struct {
	Date      time.Time `json:"date"`
	CreatedAt time.Time `json:"created_at"`
	ID        int       `json:"id"`
	HabitID   int       `json:"habit_id"`
}

type Habit struct {
	CreatedAt   time.Time `json:"created_at"`
	Title       string    `json:"title"`
	Description string    `json:"string"`
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
}

type User struct {
	CreatedAt    time.Time `json:"created_at"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	ID           int       `json:"id"`
}

type RegisterUserPayload struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginUserPayload struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CreateHabitPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"omitempty"`
}
