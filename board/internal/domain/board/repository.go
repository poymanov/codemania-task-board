package board

import "context"

type BoardRepository interface {
	Create(ctx context.Context, newBoard NewBoard) (int, error)

	GetAll(ctx context.Context, filter GetAllFilter) ([]Board, error)
}
