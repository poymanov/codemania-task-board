package v1

import (
	boardGetAllUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_all"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
)

func GetAllBoardDTOToTransport(board boardGetAllUseCase.BoardDTO) gatewayV1.GetAllBoardResponseItem {
	return gatewayV1.GetAllBoardResponseItem{
		ID:          board.Id,
		Name:        board.Name,
		Description: board.Description,
	}
}
