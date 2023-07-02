package skill

type skill struct {
	name string
	mp   int
}

func (s *skill) GetName() string {
	return s.name
}
func (s *skill) GetMp() int {
	return s.mp
}
