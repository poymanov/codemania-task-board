package board

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type GetAllBoardDTO struct {
	Id int

	Name string

	Description string
}

func ConvertTransportGetAllBoardToDTO(board *boardV1.Board) GetAllBoardDTO {
	return GetAllBoardDTO{
		Id:          int(board.Id),
		Name:        board.Name,
		Description: board.Description,
	}
}
