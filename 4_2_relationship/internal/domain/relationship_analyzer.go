package domain

type RelationShipAnalyzer interface {
	Parse(script string) RelationShipGraph
	GetMutualFriends(name1, name2 string) []string
}
