package domain

import (
	"fmt"
)

type MatchMaking struct {
	individual  Individual
	individuals []Individual
	strategy    Strategy
}

type Strategy interface {
	Sort(individuals []Individual, individual Individual) []Individual
}

func NewMatchMaking(strategy Strategy) MatchMaking {
	// 這邊我選擇直接在這邊初始化，要在外面初始化也可以，但這次是要測試不同策略，所以這個就先忽略
	return MatchMaking{
		strategy:   strategy,
		individual: NewIndividual(1, 19, 1, "", []string{"c"}, "1,35"),
		individuals: []Individual{
			NewIndividual(2, 19, 1, "", []string{"a"}, "10,10"),
			NewIndividual(3, 19, 1, "", []string{"a"}, "0,0"),
			NewIndividual(4, 19, 1, "", []string{"b"}, "5,5"),
			NewIndividual(5, 19, 1, "", []string{"a", "b"}, "1,29"),
			NewIndividual(6, 19, 1, "", []string{"c"}, "8,9"),
		},
	}
}

func (m MatchMaking) Match() {

	var strategies []string
	individuals := m.individuals
	m.strategy.Sort(m.individuals, m.individual)

	fmt.Println("The strategies are: ", strategies)
	fmt.Printf("%d match to %d\n", m.individual.Id, individuals[0].Id)
	fmt.Printf("The distance is %d\n", m.individual.GetDistance(individuals[0]))
	fmt.Printf("The habit score is %d\n", m.individual.GetHabitScore(individuals[0]))
}
