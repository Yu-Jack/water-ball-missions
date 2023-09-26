package domain

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type relationCurseSuite struct {
	suite.Suite
	mock.Mock

	rc             *RelationCurse
	cursedMocked   *MockRole
	beCursedMocked *MockRole
}

func TestAiSuite(t *testing.T) {
	suite.Run(t, new(relationCurseSuite))
}

func (s *relationCurseSuite) SetupSuite() {
	s.cursedMocked = NewMockRole(s.T())
	s.beCursedMocked = NewMockRole(s.T())

	s.rc = &RelationCurse{
		beCursed: s.beCursedMocked,
		curse:    s.cursedMocked,
	}
}

func (s *relationCurseSuite) TestGetAttackerName() {
	s.cursedMocked.On("GetNameWithTroop").Return("hi").Once()
	s.Require().Equal("hi", s.rc.GetAttackerName())
}

func (s *relationCurseSuite) TestAction() {

	tests := []struct {
		desc    string
		curseHp int
		prepare func()
	}{
		{
			desc:    "curse hp is more than zero",
			curseHp: 100,
			prepare: func() {
				s.beCursedMocked.On("GetMp").Return(50).Once()
				s.cursedMocked.On("PlusHp", 50).Return().Once()
			},
		},
		{
			desc:    "curse hp is less than zero",
			curseHp: 0,
			prepare: func() {
				s.cursedMocked.On("PlusHp", 50).Return().Times(0)
			},
		},
	}

	for _, test := range tests {
		s.cursedMocked.On("GetHp").Return(test.curseHp).Once()
		test.prepare()

		s.rc.Action()
	}

}
