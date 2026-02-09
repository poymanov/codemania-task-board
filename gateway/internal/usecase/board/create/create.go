package create

import (
	"context"

	boardGrpcClientV1 "github.com/poymanov/codemania-task-board/gateway/internal/transport/grpc/client/board/v1"
)

type UseCase struct {
	boardClient *boardGrpcClientV1.BoardClient
}

func NewUseCase(boardClient *boardGrpcClientV1.BoardClient) *UseCase {
	return &UseCase{
		boardClient: boardClient,
	}
}

func (u *UseCase) Create(ctx context.Context, dto CreateBoardDTO) (int, error) {
	createBoardRequest := boardGrpcClientV1.CreateBoardRequest{Name: dto.Name, Description: dto.Description, OwnerId: dto.OwnerId}

	boardId, err := u.boardClient.CreateBoard(ctx, createBoardRequest)
	if err != nil {
		return 0, nil
	}

	return boardId, nil
}
