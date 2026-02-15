package column

import (
	domainColumn "github.com/poymanov/codemania-task-board/board/internal/domain/column"
	boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"
)

func DomainToTransport(column domainColumn.Column) *boardV1.Column {
	return &boardV1.Column{
		Id:       int64(column.Id),
		Name:     column.Name,
		Position: float32(column.Position),
		BoardId:  int64(column.BoardId),
	}
}

func GetAllRequestToDomain(req *boardV1.ColumnServiceGetAllRequest) (domainColumn.GetAllFilter, domainColumn.GetAllSort) {
	borderId := int(req.GetFilter().GetBoardId())
	positionSort := req.GetSort().GetPosition()

	filter := domainColumn.NewGetAllFilter(borderId)
	sort := domainColumn.NewGetAllSort(positionSort)

	return filter, sort
}
