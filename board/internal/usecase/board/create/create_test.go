package create

import (
	"errors"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/mock"
)

func (s *UseCaseSuite) TestCreateError() {
	s.boardRepository.
		On("Create", s.ctx, mock.Anything).
		Return(0, errors.New(gofakeit.Word())).Once()

	res, err := s.useCase.Create(s.ctx, NewBoardDTO{})
	s.Require().Error(err)
	s.Require().Equal(0, res)
}

func (s *UseCaseSuite) TestCreateSuccess() {
	boardId := int(gofakeit.Int64())

	s.boardRepository.On("Create", s.ctx, mock.Anything).Return(boardId, nil).Once()

	res, err := s.useCase.Create(s.ctx, NewBoardDTO{})
	s.Require().NoError(err)
	s.Require().Equal(boardId, res)
}
