package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type fireFireHandler struct{}

func NewFireFireHandler() domain.SpriteHandler {
	return &fireFireHandler{}
}

func (ff *fireFireHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1.GetName() == sprite.TypeFire && c2.GetName() == sprite.TypeFire
}

func (ff *fireFireHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &worldStrategy.FailedStrategy{}
}
