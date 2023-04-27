package domain

import (
	"fmt"
	"math/rand"

	"big2/internal/domain"
)

type ai struct{}

func NewAIStrategy() domain.InputStrategy {
	return &ai{}
}

func (a ai) Input() ([]int, bool) {
	//TODO implement me
	panic("implement me")
}

func (a ai) Name() string {
	name := fmt.Sprintf("%d", rand.Intn(100000))
	fmt.Print("Hello, " + name + "!\n")
	return name
}
