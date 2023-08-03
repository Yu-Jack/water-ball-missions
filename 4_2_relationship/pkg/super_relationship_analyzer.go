package pkg

import (
	"strings"
)

type SuperRelationshipAnalyzer struct {
	data map[string][]string
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		data: make(map[string][]string),
	}
}

func (s *SuperRelationshipAnalyzer) Init(script string) {
	s.data = make(map[string][]string)

	lists := strings.Split(script, "\n")

	for _, l := range lists {
		splitStr := strings.Split(l, " -- ")
		if len(splitStr) != 2 {
			continue
		}

		host := splitStr[0]
		friend := splitStr[1]
		s.data[host] = append(s.data[host], friend)
	}
}

func (s *SuperRelationshipAnalyzer) IsMutualFriend(targetName, name1, name2 string) bool {
	name1List := s.data[name1]
	name2List := s.data[name2]

	match1, match2 := false, false

	for _, name := range name1List {
		if name == targetName {
			match1 = true
			break
		}
	}

	if !match1 {
		return false
	}

	for _, name := range name2List {
		if name == targetName {
			match2 = true
			break
		}
	}

	return match1 && match2
}
