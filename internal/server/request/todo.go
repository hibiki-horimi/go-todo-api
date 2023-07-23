package request

import (
	"time"

	"github.com/hibiki-horimi/go-todo-api/internal/domain"
	uuid "github.com/satori/go.uuid"
)

type CreateTodo struct {
	Task string `json:"task" validate:"required"`
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

type GetTodo struct {
	ID string `param:"taskId" validate:"required,uuid"`
}

func (req *GetTodo) ToTodo() *domain.Todo {
	return &domain.Todo{
		ID: uuid.Must(uuid.FromString(req.ID)),
	}
}

type UpdateTodo struct {
	ID   string `param:"taskId" validate:"required,uuid"`
	Task string `json:"task" validate:"required"`
	Done *bool  `json:"done" validate:"required"`
}

func (req *UpdateTodo) ToTodo() *domain.Todo {
	return &domain.Todo{
		ID:        uuid.Must(uuid.FromString(req.ID)),
		UpdatedAt: time.Now(),
		Task:      req.Task,
		Done:      *req.Done,
	}
}

func (req *UpdateTodo) ToTodoByID() *domain.Todo {
	return &domain.Todo{
		ID: uuid.Must(uuid.FromString(req.ID)),
	}
}
