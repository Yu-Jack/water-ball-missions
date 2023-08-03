package domain

import (
	"fmt"
	"strings"
)

type relationShipAnalyzerAdapter struct {
	superRelationshipAnalyzer *SuperRelationshipAnalyzer
	superRelationshipGraph    *SuperRelationshipGraph
	data                      map[string][]string
}

func NewRelationShipAnalyzerAdapter() RelationShipAnalyzer {
	return &relationShipAnalyzerAdapter{
		superRelationshipAnalyzer: NewSuperRelationshipAnalyzer(),
		superRelationshipGraph:    NewSuperRelationshipGraph(),
		data:                      make(map[string][]string),
	}
}

func (r *relationShipAnalyzerAdapter) Parse(script string) RelationShipGraph {
	r.superRelationshipAnalyzer.Init(r.convertScript(script))
	r.superRelationshipGraph.Init(script)
	return r.superRelationshipGraph
}

func (r *relationShipAnalyzerAdapter) GetMutualFriends(name1, name2 string) []string {
	name1List := r.data[name1]
	name2List := r.data[name2]

	friends := map[string]struct{}{}

	for _, name := range name1List {
		if r.superRelationshipAnalyzer.IsMutualFriend(name, name1, name2) {
			friends[name] = struct{}{}
		}
	}

	for _, name := range name2List {
		if r.superRelationshipAnalyzer.IsMutualFriend(name, name1, name2) {
			friends[name] = struct{}{}
		}
	}

	var keys []string
	for k := range friends {
		keys = append(keys, k)
	}
	return keys
}

func (r *relationShipAnalyzerAdapter) convertScript(script string) string {
	var output string

	lists := strings.Split(script, "\n")
	for _, l := range lists {
		splitStr := strings.Split(l, ": ")
		host := splitStr[0]
		friends := strings.Split(splitStr[1], " ")
		for _, f := range friends {
			r.data[host] = append(r.data[host], f)
			output += fmt.Sprintf("%s -- %s\n", host, f)
		}
	}

	return output
}
