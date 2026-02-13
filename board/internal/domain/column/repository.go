package column

import "context"

type ColumnRepository interface {
	Create(ctx context.Context, newColumn NewColumn) (int, error)

	GetAll(ctx context.Context, filter GetAllFilter, sort GetAllSort) ([]Column, error)
}
