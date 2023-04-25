package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type waterWaterHandler struct{}

func NewWaterWaterHandler() domain.SpriteHandler {
	return &waterWaterHandler{}
}

func (ww *waterWaterHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1.GetName() == sprite.TypeWater && c2.GetName() == sprite.TypeWater
}

func (ww *waterWaterHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &collisionStrategy.FailedCollisionStrategy{}
}
