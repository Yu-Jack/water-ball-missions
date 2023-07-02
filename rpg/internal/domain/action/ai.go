package action

import (
	"rpg/internal/domain"
)

type aiI struct {
	seed int
}

func NewAiI() domain.ActionStrategy {
	return &aiI{}
}

func (ai *aiI) S1(skillsIDs []int) int {
	return 0
}

func (ai *aiI) S2(availableIDs []int, limit int) []int {
	return []int{}
}
