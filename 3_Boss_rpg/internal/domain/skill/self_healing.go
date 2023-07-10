package skill

import (
	"fmt"

	"rpg/internal/domain"
)

type selfHealing struct {
	skill
}

func NewSelfHealing() domain.Skill {
	return &selfHealing{
		skill{mp: 50, name: domain.SkillNameSelfHealing},
	}
}

func (b selfHealing) Execute(currentRole domain.Role) {
	currentRole.PlusHp(150)
	fmt.Printf("%s 使用了 %s。\n", currentRole.GetNameWithTroop(), b.name)
}
