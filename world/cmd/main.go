package main

import (
	"world/internal/domain"
	collisionHandler "world/internal/domain/collision_handler"
)

func main() {
	handler := domain.NewHandler(
		collisionHandler.NewSpriteEmptyHandler(), domain.NewHandler(
			collisionHandler.NewWaterWaterHandler(), domain.NewHandler(
				collisionHandler.NewWaterHeroHandler(), domain.NewHandler(
					collisionHandler.NewWaterFireHandler(), domain.NewHandler(
						collisionHandler.NewHeroHeroHandler(), domain.NewHandler(
							collisionHandler.NewFireHeroHandler(), domain.NewHandler(
								collisionHandler.NewFireFireHandler(),
								nil),
						),
					),
				),
			),
		),
	)

	world := domain.NewWorld(handler)
	world.Start()
}
