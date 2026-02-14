package task

import (
	taskCreateUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/create"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

type Service struct {
	taskCreateUseCase *taskCreateUseCase.UseCase

	boardV1.UnimplementedTaskServiceServer
}

func NewService(taskCreateUseCase *taskCreateUseCase.UseCase) *Service {
	return &Service{
		taskCreateUseCase: taskCreateUseCase,
	}
}
