package v1

import boardV1 "github.com/poymanov/codemania-task-board/shared/pkg/proto/board/v1"

type BoardService struct {
	boardV1.UnimplementedBoardServiceServer
}

func NewBoardService() *BoardService {
	return &BoardService{}
}
