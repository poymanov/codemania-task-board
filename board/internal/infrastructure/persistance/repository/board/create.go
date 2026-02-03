package board

import (
	"context"
	"fmt"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

func (r *Repository) create(ctx context.Context, newBoard domainBoard.NewBoard) (int, error) {
	var id int

	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO boards (name, description, owner_id) VALUES ($1, $2, $3) RETURNING id",
		newBoard.Name, newBoard.Description, newBoard.OwnerId,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create board: %w", err)
	}

	return id, nil
}
