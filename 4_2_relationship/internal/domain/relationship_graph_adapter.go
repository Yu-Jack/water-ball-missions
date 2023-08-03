package domain

import "relationship/pkg"

type RelationShipGraphAdapter struct {
	superRelationshipGraph *pkg.SuperRelationshipGraph
}

func newRelationShipGraphAdapter() *RelationShipGraphAdapter {
	return &RelationShipGraphAdapter{
		superRelationshipGraph: pkg.NewSuperRelationshipGraph(),
	}
}

func (r *RelationShipGraphAdapter) Init(script string) {
	r.superRelationshipGraph.Init(script)
}

func (r *RelationShipGraphAdapter) HasConnection(name1, name2 string) bool {
	return r.superRelationshipGraph.HasConnection(name1, name2)
}
