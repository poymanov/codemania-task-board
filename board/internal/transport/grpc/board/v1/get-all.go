package v1

import (
	"context"

	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BoardService) GetAllByOwnerId(ctx context.Context, req *boardV1.BoardServiceGetAllByOwnerIdRequest) (*boardV1.BoardServiceGetAllByOwnerIdResponse, error) {
	boards, err := s.boardGetAllUseCase.GetAllByOwnerId(ctx, int(req.GetOwnerId()))
	if err != nil {
		log.Error().Err(err).Int("owner_id", int(req.GetOwnerId())).Msg("failed to get all boards by owner id")
		return nil, status.Errorf(codes.Internal, "error getting all boards by owner_id: %v", err)
	}

	responseBoards := make([]*boardV1.Board, 0, len(boards))

	for _, board := range boards {
		responseBoards = append(responseBoards, &boardV1.Board{
			Id:          int64(board.Id),
			Name:        board.Name,
			Description: board.Description,
			OwnerId:     int64(board.OwnerId),
		})
	}

	return &boardV1.BoardServiceGetAllByOwnerIdResponse{
		Boards: responseBoards,
	}, nil
}
