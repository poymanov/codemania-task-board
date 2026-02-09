package board

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

func (u *UseCase) Create(ctx context.Context, newBoard NewBoardDTO) (int, error) {
	nb := domainBoard.NewNewBoard(newBoard.Name, newBoard.Description, newBoard.OwnerID)

	return u.boardRepository.Create(ctx, nb)
}
