package v1

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func (c *BoardClient) CreateBoard(ctx context.Context, req CreateBoardRequest) (int, error) {
	grpcReq := &boardV1.BoardServiceCreateRequest{
		Name:        req.Name,
		Description: req.Description,
		OwnerId:     int64(req.OwnerId),
	}

	res, err := c.generatedClient.Create(ctx, grpcReq)
	if err != nil {
		return 0, err
	}

	return int(res.GetBoardId()), nil
}

func (c *BoardClient) GetAllBoard(ctx context.Context) ([]GetAllBoardDTO, error) {
	grpcReq := &boardV1.BoardServiceGetAllRequest{}

	res, err := c.generatedClient.GetAll(ctx, grpcReq)
	if err != nil {
		return []GetAllBoardDTO{}, err
	}

	responseBoards := res.GetBoards()

	dtos := make([]GetAllBoardDTO, 0, len(responseBoards))

	for _, board := range responseBoards {
		dtos = append(dtos, ConvertTransportToDTO(board))
	}

	return dtos, nil
}
