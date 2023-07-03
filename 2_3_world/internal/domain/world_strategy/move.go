package world_strategy

import "world/internal/domain"

type MoveStrategy struct{ strategy }

func (s *MoveStrategy) Action(world *domain.World, p1, p2 int) {
	world.MoveSprite(p1, p2)
}
