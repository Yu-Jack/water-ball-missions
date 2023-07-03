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
		skill{mp: 50, name: "自我治療"},
	}
}

func (b selfHealing) Execute(currentRole domain.Role) {
	currentRole.PlusHp(150)
	fmt.Printf("%s 使用了 %s。\n", currentRole.GetName(), b.name)
}
