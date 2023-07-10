package action

import "rpg/internal/domain"

type action struct {
	name domain.StrategyName
}

func (a action) GetName() domain.StrategyName {
	return a.name
}
