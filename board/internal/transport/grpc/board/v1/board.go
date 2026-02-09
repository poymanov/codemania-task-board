package v1

import (
	boardUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/board"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type BoardService struct {
	boardUseCase *boardUseCase.UseCase

	boardV1.UnimplementedBoardServiceServer
}

func NewBoardService(boardUseCase *boardUseCase.UseCase) *BoardService {
	return &BoardService{
		boardUseCase: boardUseCase,
	}
}
