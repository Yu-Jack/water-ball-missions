package world_strategy

import (
	"fmt"

	"world/internal/domain"
)

type FailedStrategy struct{ strategy }

func (s *FailedStrategy) Action(world *domain.World, p1, p2 int) {
	fmt.Println("failed")
}
