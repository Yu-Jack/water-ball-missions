package domain

type RelationShipGraph interface {
	HasConnection(name1, name2 string) bool
}
