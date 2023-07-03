package world_strategy

import "world/internal/domain"

type RemoveBothStrategy struct{ strategy }

func (s *RemoveBothStrategy) Action(world *domain.World, p1, p2 int) {
	world.RemoveSprite(p1)
	world.RemoveSprite(p2)
}
