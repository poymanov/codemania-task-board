package v1

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) Create(_ context.Context, req *boardV1.BoardServiceCreateRequest) (*boardV1.BoardServiceCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation error: %v", err)
	}

	return &boardV1.BoardServiceCreateResponse{BoardId: 1}, nil
}
