//go:generate go-enum -f=$GOFILE --nocase

package domain

type Skill interface {
	Execute(r Role)
	GetName() SkillName
	GetMp() int
}

// SkillName
/*
ENUM(
BasicAttack = "普通攻擊"
CheerUp = "鼓舞"
Curse = "詛咒"
FireBall = "火球"
WaterBall = "水球"
OnePunch = "一拳攻擊"
Petrochemical = "石化"
Poison = "下毒"
SelfExplosion = "自爆"
SelfHealing = "自我治療"
Summon = "召喚"
)
*/
type SkillName string
