package create

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
)

type UseCase struct {
	boardRepository domainBoard.BoardRepository
}

func NewUseCase(boardRepository domainBoard.BoardRepository) *UseCase {
	return &UseCase{
		boardRepository: boardRepository,
	}
}

func (u *UseCase) Create(ctx context.Context, newBoard NewBoardDTO) (int, error) {
	nb := domainBoard.NewNewBoard(newBoard.Name, newBoard.Description, newBoard.OwnerID)

	return u.boardRepository.Create(ctx, nb)
}
