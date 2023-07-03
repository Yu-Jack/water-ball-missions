package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type spriteEmptyHandler struct{}

func NewSpriteEmptyHandler() domain.SpriteHandler {
	return &spriteEmptyHandler{}
}

func (se *spriteEmptyHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1 != nil && c2 == nil
}

func (se *spriteEmptyHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &worldStrategy.MoveStrategy{}
}
