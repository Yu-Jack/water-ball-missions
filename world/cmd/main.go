package main

import (
	"world/internal/domain"
	spriteHandler "world/internal/domain/collision_handler"
)

func main() {
	handler := domain.NewHandler(
		spriteHandler.NewSpriteEmptyHandler(), domain.NewHandler(
			spriteHandler.NewWaterWaterHandler(), domain.NewHandler(
				spriteHandler.NewWaterHeroHandler(), domain.NewHandler(
					spriteHandler.NewWaterFireHandler(), domain.NewHandler(
						spriteHandler.NewHeroHeroHandler(), domain.NewHandler(
							spriteHandler.NewFireHeroHandler(), domain.NewHandler(
								spriteHandler.NewFireFireHandler(),
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
