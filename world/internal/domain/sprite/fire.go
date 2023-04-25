package sprite

type fire struct {
	sprite
}

func NewFire() Sprite {
	return &fire{
		sprite{
			name: TypeFire,
		},
	}
}
