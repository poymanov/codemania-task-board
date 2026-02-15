package v1

import (
	"context"
	"net/http"

	createColumnUseCase "github.com/poymanov/codemania-task-board/gateway/internal/usecase/column/create"
	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
	"github.com/rs/zerolog/log"
)

func (a *Api) ColumnCreate(ctx context.Context, req *gatewayV1.CreateColumnRequestBody, params gatewayV1.ColumnCreateParams) (gatewayV1.ColumnCreateRes, error) {
	createColumnDTO := createColumnUseCase.CreateColumnDTO{
		Name:    req.GetName(),
		BoardId: params.ID,
	}

	columnId, err := a.createColumnUseCase.Create(ctx, createColumnDTO)
	if err != nil {
		log.Error().Err(err).Msg("create column failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create column failed",
		}, nil
	}

	return &gatewayV1.CreateColumnResponse{
		ColumnID: columnId,
	}, nil
}

func (a *Api) ColumnDelete(ctx context.Context, params gatewayV1.ColumnDeleteParams) (gatewayV1.ColumnDeleteRes, error) {
	err := a.deleteColumnUseCase.Delete(ctx, params.ColumnId)
	if err != nil {
		log.Error().Err(err).Msg("create column failed")
		return &gatewayV1.BadRequestError{
			Code:    http.StatusBadRequest,
			Message: "Create column failed",
		}, nil
	}

	return &gatewayV1.ColumnDeleteNoContent{}, nil
}
