package strategy

import "match.making/internal/domain"

type reverse struct {
	preProcess domain.Strategy
}

func NewReverse(preProcess domain.Strategy) domain.Strategy {
	return &reverse{
		preProcess: preProcess,
	}
}

func (r reverse) Sort(individuals []domain.Individual, individual domain.Individual) []domain.Individual {
	individuals = r.preProcess.Sort(individuals, individual)

	// reverse the order of individuals
	for i, j := 0, len(individuals)-1; i < j; i, j = i+1, j-1 {
		individuals[i], individuals[j] = individuals[j], individuals[i]
	}

	return individuals
}
