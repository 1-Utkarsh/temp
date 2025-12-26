package tasks

import "time"

// Task represents the task model structure
type Task struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` //saved as enum type in the DB with values: todo, done
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// to ensure only updatable fields are modified
type UpdateTask struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` //saved as enum type in the DB with values: todo, done
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
