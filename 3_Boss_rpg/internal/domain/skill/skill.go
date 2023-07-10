package skill

import "rpg/internal/domain"

type skill struct {
	name domain.SkillName
	mp   int
}

func (s *skill) GetName() domain.SkillName {
	return s.name
}
func (s *skill) GetMp() int {
	return s.mp
}
