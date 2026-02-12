package get_all

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

func (u *UseCase) GetAllByOwnerId(ctx context.Context, ownerId int) ([]domainBoard.Board, error) {
	return u.boardRepository.GetAllByOwnerId(ctx, ownerId)
}
