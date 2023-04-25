package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
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
	return &collisionStrategy.RemoveBothCollisionStrategy{}
}
