package collision_strategy

import "world/internal/domain"

type RemoveFirstCollisionStrategy struct{ strategy }

func (s *RemoveFirstCollisionStrategy) Action(world *domain.World, p1, p2 int) {
	world.RemoveSprite(p1)
}
