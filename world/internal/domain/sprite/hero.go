package sprite

type Hero struct {
	sprite
	hp int
}

func NewHero() Sprite {
	return &Hero{
		sprite: sprite{
			name: TypeHero,
		},
		hp: 30,
	}
}

func (h *Hero) MinusHP(hp int) {
	h.hp -= hp
}

func (h *Hero) GetHP() int {
	return h.hp
}

func (h *Hero) PlusHP(hp int) {
	h.hp += hp
}
