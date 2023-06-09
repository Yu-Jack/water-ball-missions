package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type waterFireHandler struct{}

func NewWaterFireHandler() domain.SpriteHandler {
	return &waterFireHandler{}
}

func (wf *waterFireHandler) Match(c1, c2 sprite.Sprite) bool {
	return (c1.GetName() == sprite.TypeWater && c2.GetName() == sprite.TypeFire) ||
		(c1.GetName() == sprite.TypeFire && c2.GetName() == sprite.TypeWater)
}

func (wf *waterFireHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &worldStrategy.RemoveBothStrategy{}
}
