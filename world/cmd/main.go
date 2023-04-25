package main

import (
	"world/internal/domain"
	spriteHandler "world/internal/domain/sprite_handler"
)

func main() {
	handler := domain.NewHandler(
		spriteHandler.NewWaterHandler(), domain.NewHandler(
			spriteHandler.NewFireHandler(), domain.NewHandler(
				spriteHandler.NewHeroHandler(),
				nil,
			),
		),
	)

	world := domain.NewWorld(handler)
	world.Start()
}
