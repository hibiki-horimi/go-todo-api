package request

import (
	"time"

	"github.com/hibiki-horimi/go-todo-api/internal/domain"
	uuid "github.com/satori/go.uuid"
)

type CreateTodo struct {
	Task string `json:"task"`
}

func (req *CreateTodo) ToTodo() *domain.Todo {
	return &domain.Todo{
		ID:        uuid.NewV4(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Task:      req.Task,
		Done:      false,
	}
}
