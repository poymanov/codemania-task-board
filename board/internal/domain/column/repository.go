package column

import "context"

type ColumnRepository interface {
	Create(ctx context.Context, newColumn NewColumn) (int, error)
}
