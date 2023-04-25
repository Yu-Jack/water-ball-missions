package strategy

import "match.making/internal/domain"

type reverse struct {
}

func NewReverse() domain.Strategy {
	return &reverse{}
}

func (r reverse) Sort(individuals []domain.Individual, individual domain.Individual) []domain.Individual {
	// reverse the order of individuals
	for i, j := 0, len(individuals)-1; i < j; i, j = i+1, j-1 {
		individuals[i], individuals[j] = individuals[j], individuals[i]
	}

	return individuals
}
