package response

import (
	"github.com/hibiki-horimi/go-todo-api/internal/domain"
	"github.com/hibiki-horimi/go-todo-api/internal/pkg/date"
)

type Todo struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Task      string `json:"task"`
	Done      bool   `json:"done"`
}

func ToTodo(model *domain.Todo) *Todo {
	return &Todo{
		ID:        model.ID.String(),
		CreatedAt: model.CreatedAt.In(date.LocationTokyo()).Format(date.LayoutDateTime),
		UpdatedAt: model.UpdatedAt.In(date.LocationTokyo()).Format(date.LayoutDateTime),
		Task:      model.Task,
		Done:      model.Done,
	}
}

func ToTodoList(slice domain.TodoList) []*Todo {
	res := make([]*Todo, len(slice))
	for i, t := range slice {
		res[i] = ToTodo(t)
	}
	return res
}
