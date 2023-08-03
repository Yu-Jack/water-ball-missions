package domain

import (
	"fmt"
	"strings"

	"relationship/pkg"
)

type relationShipAnalyzerAdapter struct {
	superRelationshipAnalyzer *pkg.SuperRelationshipAnalyzer
	superRelationshipGraph    *pkg.SuperRelationshipGraph
	data                      map[string][]string
}

func NewRelationShipAnalyzerAdapter() RelationShipAnalyzer {
	return &relationShipAnalyzerAdapter{
		superRelationshipAnalyzer: pkg.NewSuperRelationshipAnalyzer(),
		superRelationshipGraph:    pkg.NewSuperRelationshipGraph(),
		data:                      make(map[string][]string),
	}
}

/*
Parse accepts following format string

A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z
*/
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

/*
convertScript converts following string to another format

A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z

into

A -- B
A -- C
A -- D
B -- D
B -- E
C -- E
C -- G
C -- K
C -- M
D -- K
D -- P
E -- J
E -- K
E -- L
F -- Z
*/
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
