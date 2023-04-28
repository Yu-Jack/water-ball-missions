package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type waterWaterHandler struct{}

func NewWaterWaterHandler() domain.SpriteHandler {
	return &waterWaterHandler{}
}

func (ww *waterWaterHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1.GetName() == sprite.TypeWater && c2.GetName() == sprite.TypeWater
}

func (ww *waterWaterHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &worldStrategy.FailedStrategy{}
}
