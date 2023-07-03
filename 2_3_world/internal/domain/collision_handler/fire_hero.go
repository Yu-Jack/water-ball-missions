package collision_handler

import (
	"world/internal/domain"
	"world/internal/domain/sprite"
	worldStrategy "world/internal/domain/world_strategy"
)

type fireHeroHandler struct{}

func NewFireHeroHandler() domain.SpriteHandler {
	return &fireHeroHandler{}
}

func (fh *fireHeroHandler) Match(c1, c2 sprite.Sprite) bool {
	return (c1.GetName() == sprite.TypeFire && c2.GetName() == sprite.TypeHero) ||
		(c1.GetName() == sprite.TypeHero && c2.GetName() == sprite.TypeFire)
}

func (fh *fireHeroHandler) Collision(c1, c2 sprite.Sprite) domain.CollisionStrategy {

	// c1 是 fire
	if c1.GetName() == sprite.TypeFire {
		hero := c2.(*sprite.Hero)
		hero.MinusHP(10)
		if hero.GetHP() <= 0 {
			return &worldStrategy.RemoveBothStrategy{}
		}

		return &worldStrategy.RemoveFirstStrategy{}
	}

	// c1 是 hero
	if c1.GetName() == sprite.TypeHero {
		hero := c1.(*sprite.Hero)
		hero.MinusHP(10)

		if hero.GetHP() <= 0 {
			return &worldStrategy.RemoveBothStrategy{}
		}

		return &worldStrategy.MoveStrategy{}
	}

	return &worldStrategy.FailedStrategy{}
}
