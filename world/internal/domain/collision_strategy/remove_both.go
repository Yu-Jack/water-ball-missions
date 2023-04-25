package collision_strategy

import "world/internal/domain"

type RemoveBothCollisionStrategy struct{ collisionStrategy }

func (s *RemoveBothCollisionStrategy) Action(world *domain.World, p1, p2 int) {
	world.RemoveSprite(p1)
	world.RemoveSprite(p2)
}
