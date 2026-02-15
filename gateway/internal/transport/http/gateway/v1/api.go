package v1

import (
	"context"
	"net/http"

	createBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	getAllBoardUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_all"
	createColumnUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/create"
	deleteColumnUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/delete"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
)

type Api struct {
	createBoardUseCase  *createBoardUseCase.UseCase
	getAllBoardUseCase  *getAllBoardUseCase.UseCase
	createColumnUseCase *createColumnUseCase.UseCase
	deleteColumnUseCase *deleteColumnUseCase.UseCase
}

func NewApi(
	createBoardUseCase *createBoardUseCase.UseCase,
	getAllBoardUseCase *getAllBoardUseCase.UseCase,
	createColumnUseCase *createColumnUseCase.UseCase,
	deleteColumnUseCase *deleteColumnUseCase.UseCase,
) *Api {
	return &Api{
		createBoardUseCase:  createBoardUseCase,
		getAllBoardUseCase:  getAllBoardUseCase,
		createColumnUseCase: createColumnUseCase,
		deleteColumnUseCase: deleteColumnUseCase,
	}
}

func (a *Api) NewError(_ context.Context, err error) *gatewayV1.GenericErrorStatusCode {
	return &gatewayV1.GenericErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: gatewayV1.GenericError{
			Code:    gatewayV1.NewOptInt(http.StatusInternalServerError),
			Message: gatewayV1.NewOptString(err.Error()),
		},
	}
}
