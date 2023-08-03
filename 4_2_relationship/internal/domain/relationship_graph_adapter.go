package domain

type relationShipGraphAdapter struct{}

func newRelationShipGraphAdapter() RelationShipGraph {
	return &relationShipGraphAdapter{}
}

func (r relationShipGraphAdapter) HasConnection(name1, name2 string) bool {
	//TODO implement me
	panic("implement me")
}
