package action

import (
	"rpg/internal/domain"
)

type aiI struct {
	action
	seed int
}

func NewAiI() domain.ActionStrategy {
	return &aiI{
		seed: 0,
		action: action{
			name: domain.StrategyNameAI,
		},
	}
}

func (ai *aiI) S1(skillsIDs []int) int {
	target := ai.seed % len(skillsIDs)
	ai.seed++
	return target
}

func (ai *aiI) S2(availableIDs []int, limit int, list string) []int {
	if len(availableIDs) == limit {
		return availableIDs
	}

	var targets []int

	for i := 0; i < limit; i++ {
		targets = append(targets, (ai.seed+i)%len(availableIDs))
	}

	ai.seed++
	return targets
}
