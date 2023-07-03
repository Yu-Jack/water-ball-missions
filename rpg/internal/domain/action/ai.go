package action

import (
	"fmt"

	"rpg/internal/domain"
)

type aiI struct {
	seed int
}

func NewAiI() domain.ActionStrategy {
	return &aiI{seed: 0}
}

func (ai *aiI) S1(skillsIDs []int) int {
	fmt.Println(ai.seed)
	target := ai.seed % len(skillsIDs)
	ai.seed++
	return target
}

func (ai *aiI) S2(availableIDs []int, limit int) []int {
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
