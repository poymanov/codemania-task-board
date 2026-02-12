package board

import "context"

type BoardRepository interface {
	Create(ctx context.Context, newBoard NewBoard) (int, error)

	GetAllByOwnerId(ctx context.Context, ownerId int) ([]Board, error)
}
