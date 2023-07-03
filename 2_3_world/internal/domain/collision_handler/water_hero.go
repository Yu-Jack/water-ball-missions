package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type waterHeroHandler struct{}

func NewWaterHeroHandler() domain.SpriteHandler {
	return &waterHeroHandler{}
}

func (wh *waterHeroHandler) Match(c1, c2 sprite.Sprite) bool {
	return (c1.GetName() == sprite.TypeWater && c2.GetName() == sprite.TypeHero) ||
		(c1.GetName() == sprite.TypeHero && c2.GetName() == sprite.TypeWater)
}

func (wh *waterHeroHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {

	// c1 是 water
	if c1.GetName() == sprite.TypeWater {
		c2.(*sprite.Hero).PlusHP(10)
		return &worldStrategy.RemoveFirstStrategy{}
	}

	// c1 是 hero
	if c1.GetName() == sprite.TypeHero {
		c1.(*sprite.Hero).PlusHP(10)
		return &worldStrategy.MoveStrategy{}
	}

	return &worldStrategy.FailedStrategy{}
}
