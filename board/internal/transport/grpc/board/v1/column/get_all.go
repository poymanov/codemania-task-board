package column

import (
	"context"

	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetAll(ctx context.Context, req *boardV1.ColumnServiceGetAllRequest) (*boardV1.ColumnServiceGetAllResponse, error) {
	borderId := int(req.GetFilter().GetBoardId())
	positionSort := req.GetSort().GetPosition()

	filter := domainColumn.NewGetAllFilter(borderId)
	sort := domainColumn.NewGetAllSort(positionSort)

	columns, err := s.columnGetAllUseCase.GetAll(ctx, filter, sort)
	if err != nil {
		log.Error().Err(err).Any("req", req).Msg("failed to get all boards")
		return nil, status.Error(codes.Internal, "failed to get all boards")
	}

	responseColumns := make([]*boardV1.Column, 0, len(columns))

	for _, column := range columns {
		responseColumns = append(responseColumns, &boardV1.Column{
			Id:       int64(column.Id),
			Name:     column.Name,
			Position: float32(column.Position),
			BoardId:  int64(column.BoardId),
		})
	}

	return &boardV1.ColumnServiceGetAllResponse{
		Columns: responseColumns,
	}, nil
}
