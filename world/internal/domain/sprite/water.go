package sprite

type water struct {
	sprite
}

func NewWater() Sprite {
	return &water{
		sprite{
			name: TypeWater,
		},
	}
}
