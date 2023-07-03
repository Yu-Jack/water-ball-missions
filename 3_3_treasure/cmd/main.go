package main

import (
	"math/rand"
	"time"

	"treasure/internal/domain"
	"treasure/internal/domain/state"
	"treasure/internal/domain/treasure"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	m := domain.NewMap()

	m.SetRole(domain.NewRoleMonster(m.RandomEmptyPosition(), state.NewStateNormal()))

	m.SetUpMonsterGenerator(func() {
		if rand.Int()%2 == 0 {
			m.SetRole(domain.NewRoleMonster(m.RandomEmptyPosition(), state.NewStateNormal()))
		}
	})

	m.SetUpTreasureGenerator(func() {
		m.SetTreasure(treasure.GenerateRandomTreasure(m.RandomEmptyPosition()))
	})

	// character
	m.SetRole(domain.NewRoleCharacter(m.RandomEmptyPosition(), state.NewStateNormal()))

	// obstacles
	for i := 0; i < 5; i++ {
		m.SetObstacle(domain.Obstacle{Position: m.RandomEmptyPosition()})
	}

	m.GenerateMonster()

	m.Start()
}
