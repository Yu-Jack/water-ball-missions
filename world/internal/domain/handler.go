package domain

import (
	"fmt"

	"world/internal/domain/sprite"
)

type SpriteHandler interface {
	Match(c1, c2 sprite.Sprite) bool
	Collision(c1, c2 sprite.Sprite) CollisionStrategy
}

type CollisionStrategy interface {
	Action(world *World, p1, p2 int)
}

type Handler interface {
	Collision(c1, c2 sprite.Sprite) CollisionStrategy // 0: nothing, 1: failed, 2: remove
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

func (h *handler) Collision(c1, c2 sprite.Sprite) CollisionStrategy {
	if c1 == nil {
		return nil
	}

	fmt.Printf("%#v\n", h.spriteHandler)

	if h.spriteHandler.Match(c1, c2) {
		return h.spriteHandler.Collision(c1, c2)
	}

	if h.next != nil {
		return h.next.Collision(c1, c2)
	}

	return nil
}
