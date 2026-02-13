package column

import (
	columnCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/create"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type ColumnService struct {
	columnCreateUseCase *columnCreateUseCase.UseCase

	boardV1.UnimplementedColumnServiceServer
}

func NewColumnService(columnCreateUseCase *columnCreateUseCase.UseCase) *ColumnService {
	return &ColumnService{
		columnCreateUseCase: columnCreateUseCase,
	}
}
