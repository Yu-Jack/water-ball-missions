package world_strategy

import "world/internal/domain"

type RemoveFirstStrategy struct{ strategy }

func (s *RemoveFirstStrategy) Action(world *domain.World, p1, p2 int) {
	world.RemoveSprite(p1)
}
