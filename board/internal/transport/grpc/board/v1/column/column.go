package column

import (
	columnCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/create"
	columnDeleteUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/delete"
	columnGetAllUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/column/get_all"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type ColumnService struct {
	columnCreateUseCase *columnCreateUseCase.UseCase
	columnGetAllUseCase *columnGetAllUseCase.UseCase
	columnDeleteUseCase *columnDeleteUseCase.UseCase

	boardV1.UnimplementedColumnServiceServer
}

func NewColumnService(columnCreateUseCase *columnCreateUseCase.UseCase, columnGetAllUseCase *columnGetAllUseCase.UseCase, columnDeleteUseCase *columnDeleteUseCase.UseCase) *ColumnService {
	return &ColumnService{
		columnCreateUseCase: columnCreateUseCase,
		columnGetAllUseCase: columnGetAllUseCase,
		columnDeleteUseCase: columnDeleteUseCase,
	}
}
