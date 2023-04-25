package collision_strategy

import "world/internal/domain"

type MoveCollisionStrategy struct{ strategy }

func (s *MoveCollisionStrategy) Action(world *domain.World, p1, p2 int) {
	world.MoveSprite(p1, p2)
}
