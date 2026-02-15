package v1

import (
	"context"
	"net/http"

	boardCreateUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/create"
	boardGetAllUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/board/get_all"
	columnCreateUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/create"
	columnDeleteUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/delete"
	columnUpdatePositionUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/update_position"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
)

type Api struct {
	boardCreateUseCase          *boardCreateUseCase.UseCase
	boardGetAllUseCase          *boardGetAllUseCase.UseCase
	columnCreateUseCase         *columnCreateUseCase.UseCase
	columnDeleteUseCase         *columnDeleteUseCase.UseCase
	columnUpdatePositionUseCase *columnUpdatePositionUseCase.UseCase
}

func NewApi(
	boardCreateUseCase *boardCreateUseCase.UseCase,
	boardGetAllUseCase *boardGetAllUseCase.UseCase,
	columnCreateUseCase *columnCreateUseCase.UseCase,
	columnDeleteUseCase *columnDeleteUseCase.UseCase,
	columnUpdatePositionUseCase *columnUpdatePositionUseCase.UseCase,
) *Api {
	return &Api{
		boardCreateUseCase:          boardCreateUseCase,
		boardGetAllUseCase:          boardGetAllUseCase,
		columnCreateUseCase:         columnCreateUseCase,
		columnDeleteUseCase:         columnDeleteUseCase,
		columnUpdatePositionUseCase: columnUpdatePositionUseCase,
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
