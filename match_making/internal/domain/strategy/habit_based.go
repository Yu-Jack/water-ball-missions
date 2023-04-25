package strategy

import (
	"sort"

	"match.making/internal/domain"
)

type habitBased struct {
}

func NewHabitBased() domain.Strategy {
	return &habitBased{}
}

func (d habitBased) Sort(individuals []domain.Individual, individual domain.Individual) []domain.Individual {
	sort.Slice(individuals, func(i, j int) bool {
		return individuals[i].GetHabitScore(individual) > individuals[i].GetHabitScore(individuals[j])
	})

	return individuals
}
