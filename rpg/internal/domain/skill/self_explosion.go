package skill

import (
	"fmt"
	"strings"

	"rpg/internal/domain"
)

type selfExplosion struct {
	skill
}

func NewSelfExplosion() domain.Skill {
	return &selfExplosion{
		skill{mp: 200, name: "自爆"},
	}
}

func (b selfExplosion) Execute(currentRole domain.Role) {
	allRoles := currentRole.GetRPG().GetAllRolesExcludeSelf(currentRole.GetTroopID(), currentRole.GetID())
	damage := 150 + currentRole.GetExtraStr()

	var output []string
	for _, targetRole := range allRoles {
		output = append(output, targetRole.GetNameWithTroop())

	}
	fmt.Printf("%s 對 %s 使用了 %s。\n", currentRole.GetNameWithTroop(), strings.Join(output, ", "), b.name)

	for _, targetRole := range allRoles {
		domain.LogDamage(currentRole.GetNameWithTroop(), targetRole.GetNameWithTroop(), damage)
		targetRole.MinusHp(damage)
	}

	currentRole.MinusHp(currentRole.GetHp())
}
