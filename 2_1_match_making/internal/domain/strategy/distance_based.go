package strategy

import (
	"sort"

	"match.making/internal/domain"
)

type distanceBased struct {
}

func NewDistanceBased() domain.Strategy {
	return &distanceBased{}
}

func (d distanceBased) Sort(individuals []domain.Individual, individual domain.Individual) []domain.Individual {
	sort.Slice(individuals, func(i, j int) bool {
		return individuals[i].GetDistance(individual) > individuals[j].GetDistance(individual)
	})

	return individuals
}
