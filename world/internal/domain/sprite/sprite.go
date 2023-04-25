package sprite

type Sprite interface {
	GetName() Type
}

type sprite struct {
	name Type
}

func (s *sprite) GetName() Type { return s.name }
