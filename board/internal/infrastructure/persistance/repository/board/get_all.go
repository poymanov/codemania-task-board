package board

import (
	"context"

	"github.com/jackc/pgx/v5"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

func (r *Repository) GetAllByOwnerId(ctx context.Context, ownerId int) ([]domainBoard.Board, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM boards WHERE owner_id = $1", ownerId)
	if err != nil {
		return []domainBoard.Board{}, err
	}

	defer rows.Close()

	boardModels, err := pgx.CollectRows(rows, pgx.RowToStructByName[Board])
	if err != nil {
		return []domainBoard.Board{}, err
	}

	boards := make([]domainBoard.Board, 0, len(boardModels))

	for _, model := range boardModels {
		board := domainBoard.Board{
			Id:          model.Id,
			Name:        model.Name,
			Description: model.Description,
			OwnerId:     ownerId,
		}

		boards = append(boards, board)
	}

	return boards, nil
}
