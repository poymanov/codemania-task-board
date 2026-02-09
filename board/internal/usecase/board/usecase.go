package board

import domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"

type NewBoardDTO struct {
	Name string

	Description string

	OwnerID int
}

type UseCase struct {
	boardRepository domainBoard.BoardRepository
}

func NewUseCase(boardRepository domainBoard.BoardRepository) *UseCase {
	return &UseCase{
		boardRepository: boardRepository,
	}
}
