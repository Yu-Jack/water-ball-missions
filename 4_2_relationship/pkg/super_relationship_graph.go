package pkg

import (
	"strings"
)

type SuperRelationshipGraph struct {
	data map[string][]string
}

func NewSuperRelationshipGraph() *SuperRelationshipGraph {
	return &SuperRelationshipGraph{
		data: make(map[string][]string),
	}
}

func (s *SuperRelationshipGraph) Init(script string) {
	s.data = make(map[string][]string)

	lists := strings.Split(script, "\n")
	for _, l := range lists {
		splitStr := strings.Split(l, ": ")
		host := splitStr[0]
		friends := strings.Split(splitStr[1], " ")
		for _, f := range friends {
			s.data[host] = append(s.data[host], f)
		}
	}
}

func (s *SuperRelationshipGraph) HasConnection(name1, name2 string) bool {
	q := []string{name1} // start point
	visited := make(map[string]struct{})

	for len(q) > 0 {
		pop := q[0]
		q = q[1:]

		if pop == name2 {
			return true
		}

		if _, ok := visited[pop]; ok {
			continue
		}
		visited[pop] = struct{}{}

		for _, friend := range s.data[pop] {
			q = append(q, friend)
		}
	}

	return false
}
