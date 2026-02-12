package get_all

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	domainBoard "github.com/poymanov/codemania-task-board/board/internal/domain/board"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestGetAllByOwnerError() {
	s.boardRepository.
		On("GetAllByOwnerId", s.ctx, mock.Anything).
		Return([]domainBoard.Board{}, errors.New(gofakeit.Word())).Once()

	res, err := s.useCase.GetAllByOwnerId(s.ctx, int(gofakeit.Int64()))
	s.Require().Error(err)
	s.Require().Empty(res)
}

func (s *UseCaseSuite) TestGetAllByOwnerSuccess() {
	ownerId := int(gofakeit.Int64())
	s.boardRepository.On("GetAllByOwnerId", s.ctx, mock.Anything).Return([]domainBoard.Board{{OwnerId: ownerId}}, nil).Once()

	res, err := s.useCase.GetAllByOwnerId(s.ctx, ownerId)
	s.Require().NoError(err)
	s.Require().NotEmpty(res)
}
