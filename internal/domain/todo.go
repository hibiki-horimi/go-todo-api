package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
	Task      string    `gorm:"column:task;not null"`
	Done      bool      `gorm:"column:done;not null"`
}

type TodoList []*Todo

func (Todo) TableName() string {
	return "todo"
}
