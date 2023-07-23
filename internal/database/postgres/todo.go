package postgres

import (
	"context"

	"github.com/hibiki-horimi/go-todo-api/internal/domain"
)

type Todo interface {
	Get(ctx context.Context, condition *domain.Todo) (*domain.Todo, error)
	Find(ctx context.Context, condition *domain.Todo) (domain.TodoList, error)
	Create(ctx context.Context, model *domain.Todo) error
	Update(ctx context.Context, model *domain.Todo) error
	Delete(ctx context.Context, condition *domain.Todo) error
}

type todo struct{}

func (pg *todo) Get(ctx context.Context, condition *domain.Todo) (*domain.Todo, error) {
	var result *domain.Todo
	if err := DBFromContext(ctx).First(&result, condition).Error; err != nil {
		return nil, err
	}
	return result, nil
}

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

func (pg *todo) Update(ctx context.Context, model *domain.Todo) error {
	columns := []string{
		"updated_at",
		"task",
		"done",
	}
	if err := DBFromContext(ctx).Select(columns).Updates(model).Error; err != nil {
		return err
	}
	return nil
}

func (pg *todo) Delete(ctx context.Context, condition *domain.Todo) error {
	if err := DBFromContext(ctx).Delete(condition).Error; err != nil {
		return err
	}
	return nil
}
