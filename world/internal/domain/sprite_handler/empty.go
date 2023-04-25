package sprite_handler

import (
	"world/internal/domain"
	collisionStrategy "world/internal/domain/collision_strategy"
	"world/internal/domain/sprite"
)

type spriteEmptyHandler struct{}

func NewSpriteEmptyHandler() domain.SpriteHandler {
	return &spriteEmptyHandler{}
}

func (se *spriteEmptyHandler) Match(c1, c2 sprite.Sprite) bool {
	return c1 != nil && c2 == nil
}

func (se *spriteEmptyHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {
	return &collisionStrategy.MoveCollisionStrategy{}
}
