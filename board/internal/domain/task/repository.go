package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, newTask NewTask) (int, error)

	GetAll(ctx context.Context, filter GetAllFilter, sort GetAllSort) ([]Task, error)

	Delete(ctx context.Context, id int) error
}
