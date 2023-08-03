package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestRelationShipAnalyzerAdapterSuit(t *testing.T) {
	suite.Run(t, new(relationShipAnalyzerAdapterSuit))
}

type relationShipAnalyzerAdapterSuit struct {
	suite.Suite
	analyzer RelationShipAnalyzer
}

func (s *relationShipAnalyzerAdapterSuit) SetupSuite() {
	s.analyzer = NewRelationShipAnalyzerAdapter()
}

func (s *relationShipAnalyzerAdapterSuit) TestGetMutualFriends() {
	s.analyzer.Parse(`A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z`)

	tests := []struct {
		name1   string
		name2   string
		friends []string
	}{
		{
			name1:   "A",
			name2:   "B",
			friends: []string{"D"},
		},
		{
			name1:   "A",
			name2:   "D",
			friends: []string{"B"},
		},
		{
			name1:   "D",
			name2:   "E",
			friends: []string{"B", "K"},
		},
	}

	for _, test := range tests {
		friends := s.analyzer.GetMutualFriends(test.name1, test.name2)
		s.Require().Equal(test.friends, friends)
	}
}

func (s *relationShipAnalyzerAdapterSuit) TestGraph() {
	graph := s.analyzer.Parse(`A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z`)

	tests := []struct {
		name1      string
		name2      string
		connection bool
	}{
		{
			name1:      "A",
			name2:      "B",
			connection: true,
		},
		{
			name1:      "A",
			name2:      "D",
			connection: true,
		},
		{
			name1:      "A",
			name2:      "F",
			connection: false,
		},
	}

	for _, test := range tests {
		connection := graph.HasConnection(test.name1, test.name2)
		s.Require().Equal(test.connection, connection)
	}
}
