package v1

import (
	"context"

	gatewayV1 "github.com/poymanov/codemania-task-board/shared/pkg/openapi/gateway/v1"
)

func (a *Api) BoardCreate(_ context.Context, req *gatewayV1.CreateBoardRequestBody) (gatewayV1.BoardCreateRes, error) {
	return &gatewayV1.CreateBoardResponse{
		BoardID: 1,
	}, nil
}
