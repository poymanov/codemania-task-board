package task

import (
	"context"
)

type TaskRepository interface {
	Create(ctx context.Context, newTask NewTask) (int, error)
}
