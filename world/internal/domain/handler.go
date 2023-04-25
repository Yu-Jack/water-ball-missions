package domain

import "world/internal/domain/sprite"

type SpriteHandler interface {
	Match(sprite sprite.Sprite) bool
	Collision(spriteX, spriteY sprite.Sprite) CollisionStrategy
}

type Handler interface {
	Collision(spriteX, spriteY sprite.Sprite) CollisionStrategy // 0: nothing, 1: failed, 2: remove
}

type handler struct {
	world         *World
	spriteHandler SpriteHandler
	next          Handler
}

func NewHandler(spriteHandler SpriteHandler, next Handler) Handler {
	return &handler{
		spriteHandler: spriteHandler,
		next:          next,
	}
}

func (h *handler) Collision(spriteX, spriteY sprite.Sprite) CollisionStrategy {
	if spriteX == nil {
		return &FailedCollisionStrategy{}
	}

	if h.spriteHandler.Match(spriteX) {
		return h.spriteHandler.Collision(spriteX, spriteY)
	}

	if h.next != nil {
		return h.next.Collision(spriteX, spriteY)
	}

	return &FailedCollisionStrategy{}
}
