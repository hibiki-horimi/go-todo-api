package postgres

import (
	"context"

	"github.com/hibiki-horimi/go-todo-api/internal/domain"
)

type Todo interface {
	Find(ctx context.Context, condition *domain.Todo) (domain.TodoList, error)
	Create(ctx context.Context, model *domain.Todo) error
}

type todo struct{}

func (pg *todo) Find(ctx context.Context, condition *domain.Todo) (domain.TodoList, error) {
	var result domain.TodoList
	if err := DBFromContext(ctx).Find(&result, condition).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (pg *todo) Create(ctx context.Context, model *domain.Todo) error {
	if err := DBFromContext(ctx).Create(model).Error; err != nil {
		return err
	}
	return nil
}
