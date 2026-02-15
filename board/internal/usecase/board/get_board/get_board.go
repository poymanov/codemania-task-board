package get_board

import (
	"context"

	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	domainCommon "github.com/poymanov/codemania-task-board/board/internal/domain/common"
	domainTask "github.com/poymanov/codemania-task-board/board/internal/domain/task"
)

type UseCase struct {
	boardRepository  domainBoard.BoardRepository
	columnRepository domainColumn.ColumnRepository
	taskRepository   domainTask.TaskRepository
}

func NewUseCase(
	boardRepository domainBoard.BoardRepository,
	columnRepository domainColumn.ColumnRepository,
	taskRepository domainTask.TaskRepository,
) *UseCase {
	return &UseCase{
		boardRepository:  boardRepository,
		columnRepository: columnRepository,
		taskRepository:   taskRepository,
	}
}

func (u *UseCase) GetBoard(ctx context.Context, id int) (domainCommon.Board, error) {
	board, err := u.boardRepository.GetById(ctx, id)
	if err != nil {
		return domainCommon.Board{}, err
	}

	boardFilter := domainColumn.NewGetAllFilter(board.Id)
	boardSort := domainColumn.NewGetAllSort("asc")

	columns, err := u.columnRepository.GetAll(ctx, boardFilter, boardSort)
	if err != nil {
		return domainCommon.Board{}, nil
	}

	commonColumns := make([]domainCommon.Column, 0, len(columns))

	for _, column := range columns {
		taskFilter := domainTask.NewGetAllFilter(column.Id)
		taskSort := domainTask.NewGetAllSort("asc")

		tasks, errTasks := u.taskRepository.GetAll(ctx, taskFilter, taskSort)

		if errTasks != nil {
			return domainCommon.Board{}, errTasks
		}

		commonTasks := make([]domainCommon.Task, 0, len(tasks))

		for _, task := range tasks {
			commonTasks = append(commonTasks, domainCommon.Task{
				Id:          task.Id,
				Description: task.Description,
				Assignee:    task.Assignee,
				Position:    task.Position,
			})
		}

		commonColumns = append(commonColumns, domainCommon.Column{
			Id:       column.Id,
			Name:     column.Name,
			Position: column.Position,
			Tasks:    commonTasks,
		})
	}

	commonBoard := domainCommon.Board{
		Id:          board.Id,
		Name:        board.Name,
		Description: board.Description,
		OwnerId:     board.OwnerId,
		Columns:     commonColumns,
	}

	return commonBoard, err
}
