package domain

import "fmt"

type role struct {
	name           string
	troopID        int
	ID             int
	hp             int
	mp             int
	str            int
	extraStr       int
	state          State
	skills         []Skill
	actionAble     bool
	actionStrategy ActionStrategy
	rpg            RPG
}

type Role interface {
	PlusHp(hp int)
	MinusHp(hp int)
	MinusMp(mp int)
	SetExtraStr(str int)
	SetActionAble(bool)
	GetState() State
	SetState(newState State)
	ActionS1(skillIDs []int) int
	ActionS2(availableIDs []int, limit int) []int
	GetTroopID() int
	GetStr() int
	GetExtraStr() int
	TakeAction()
	GetRPG() RPG
	GetName() string
	GetHp() int
}

func (r *role) GetTroopID() int               { return r.troopID }
func (r *role) GetRPG() RPG                   { return r.rpg }
func (r *role) GetStr() int                   { return r.str }
func (r *role) GetExtraStr() int              { return r.extraStr }
func (r *role) SetExtraStr(extraStr int)      { r.extraStr = extraStr }
func (r *role) SetActionAble(actionAble bool) { r.actionAble = actionAble }
func (r *role) GetHp() int                    { return r.hp }
func (r *role) GetState() State               { return r.state }

func (r *role) SetState(newState State) {
	if newState.GetName() != r.state.GetName() {
		r.state.Exit()
	}

	newState.SetRole(r)
	r.state = newState
	r.state.Enter()
}

func (r *role) ActionS1(skillIDs []int) int {
	return r.actionStrategy.S1(skillIDs)
}

func (r *role) ActionS2(availableIDs []int, limit int) []int {
	return r.actionStrategy.S2(availableIDs, limit)
}

func (r *role) PlusHp(hp int) {
	r.hp += hp
}

func (r *role) MinusHp(hp int) {
	r.hp -= hp

	if r.hp <= 0 {
		fmt.Printf("%s 死亡。\n", r.GetName())
	}
}

func (r *role) MinusMp(mp int) {
	r.mp -= mp
}

func (r *role) TakeAction() {
	fmt.Printf(
		"輪到 %s (%s)。\n",
		r.GetName(), r.GetStatus(),
	)

	if !r.actionAble {
		fmt.Printf("%s 無法行動。\n", r.GetName())
		return
	}

	var skillIDs []int
	for i := range r.skills {
		skillIDs = append(skillIDs, i)
	}

	for {
		fmt.Printf(
			"選擇行動：%s\n",
			r.GetAllSkillName(),
		)

		targetSkillID := r.ActionS1(skillIDs)
		targetSkill := r.skills[targetSkillID]

		if targetSkill.GetMp() > r.mp {
			fmt.Println("你缺乏 MP，不能進行此行動。")
			continue
		}

		targetSkill.Execute(r)
		r.MinusMp(targetSkill.GetMp())
		break
	}
}

func (r *role) GetName() string {
	return fmt.Sprintf("[%d]%s", r.troopID, r.name)
}

func (r *role) GetStatus() string {
	return fmt.Sprintf(
		"HP: %d, MP: %d, STR: %d, State: %s",
		r.hp, r.mp, r.str, r.state.GetName(),
	)
}

func (r *role) GetAllSkillName() string {
	str := ""

	for i, skill := range r.skills {
		str += fmt.Sprintf("(%d) %s ", i, skill.GetName())
	}

	return str
}
