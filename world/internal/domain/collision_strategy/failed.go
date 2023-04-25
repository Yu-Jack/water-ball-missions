package collision_strategy

import (
	"fmt"

	"world/internal/domain"
)

type FailedCollisionStrategy struct{ collisionStrategy }

func (s *FailedCollisionStrategy) Action(world *domain.World, p1, p2 int) {
	fmt.Println("failed")
}
