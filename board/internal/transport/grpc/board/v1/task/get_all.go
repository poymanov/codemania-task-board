package task

import (
	"context"

	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAll(ctx context.Context, req *boardV1.TaskServiceGetAllRequest) (*boardV1.TaskServiceGetAllResponse, error) {
	columnId := int(req.GetFilter().GetColumnId())
	positionSort := req.GetSort().GetPosition()

	filter := domainTask.NewGetAllFilter(columnId)
	sort := domainTask.NewGetAllSort(positionSort)

	tasks, err := s.taskGetAllUseCase.GetAll(ctx, filter, sort)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to get all tasks")
		return nil, status.Error(codes.Internal, "failed to get all tasks")
	}

	responseTasks := make([]*boardV1.Task, 0, len(tasks))

	for _, task := range tasks {
		responseTasks = append(responseTasks, &boardV1.Task{
			Id:          int64(task.Id),
			Title:       task.Title,
			Description: task.Description,
			Assignee:    task.Assignee,
			Position:    float32(task.Position),
			ColumnId:    int64(task.ColumnId),
		})
	}

	return &boardV1.TaskServiceGetAllResponse{
		Tasks: responseTasks,
	}, nil
}
