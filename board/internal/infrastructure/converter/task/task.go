package task

import (
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
	taskUpdatePositionUseCase "github.com/poymanov/codemania-task-board/board/internal/usecase/task/update_position"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func DomainToTransport(task domainTask.Task) *boardV1.Task {
	return &boardV1.Task{
		Id:          int64(task.Id),
		Title:       task.Title,
		Description: task.Description,
		Assignee:    task.Assignee,
		Position:    float32(task.Position),
		ColumnId:    int64(task.ColumnId),
	}
}

func GetAllRequestToDomain(req *boardV1.TaskServiceGetAllRequest) (domainTask.GetAllFilter, domainTask.GetAllSort) {
	columnId := int(req.GetFilter().GetColumnId())
	positionSort := req.GetSort().GetPosition()

	filter := domainTask.NewGetAllFilter(columnId)
	sort := domainTask.NewGetAllSort(positionSort)

	return filter, sort
}

func UpdatePositionRequestToUseCaseDTO(req *boardV1.TaskServiceUpdatePositionRequest) taskUpdatePositionUseCase.UpdatePositionDTO {
	return taskUpdatePositionUseCase.UpdatePositionDTO{
		LeftPosition:  float64(req.GetLeftPosition()),
		RightPosition: float64(req.GetRightPosition()),
	}
}
