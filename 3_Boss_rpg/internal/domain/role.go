package domain

import (
	"fmt"
	"strings"
)

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
	relations      []RelationObserver
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
	ActionS2(availableIDs []int, limit int, list string) []int
	GetTroopID() int
	GetID() int
	GetStr() int
	GetExtraStr() int
	TakeAction()
	GetRPG() RPG
	GetNameWithTroop() string
	GetName() string
	GetHp() int
	GetMp() int
	AddRelation(observer RelationObserver) bool
	GetRelations() []RelationObserver
}

func NewRole(name string, hp, mp, str int, state State, skills []Skill, strategy ActionStrategy) Role {
	// `rpg` && `ID` && `troopID` will be set when it's assigned into troop
	role := &role{
		name:           name,
		hp:             hp,
		mp:             mp,
		str:            str,
		extraStr:       0,
		skills:         skills,
		state:          state,
		actionAble:     true,
		actionStrategy: strategy,
	}

	role.SetState(state)
	return role
}

func (r *role) GetID() int                    { return r.ID }
func (r *role) GetTroopID() int               { return r.troopID }
func (r *role) GetRPG() RPG                   { return r.rpg }
func (r *role) GetStr() int                   { return r.str }
func (r *role) GetExtraStr() int              { return r.extraStr }
func (r *role) SetExtraStr(extraStr int)      { r.extraStr = extraStr }
func (r *role) SetActionAble(actionAble bool) { r.actionAble = actionAble }
func (r *role) GetHp() int                    { return r.hp }
func (r *role) GetMp() int                    { return r.mp }
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

func (r *role) ActionS2(availableIDs []int, limit int, list string) []int {
	return r.actionStrategy.S2(availableIDs, limit, list)
}

func (r *role) PlusHp(hp int) {
	r.hp += hp
}

func (r *role) MinusHp(hp int) {
	r.hp -= hp

	if r.hp <= 0 {
		fmt.Printf("%s 死亡。\n", r.GetNameWithTroop())
		r.state.Die()
	}
}

func (r *role) MinusMp(mp int) {
	r.mp -= mp
}

func (r *role) PrintInformation() {
	fmt.Printf(
		"輪到 %s (%s)。\n",
		r.GetNameWithTroop(), r.GetStatus(),
	)
}

func (r *role) TakeAction() {
	if !r.actionAble {
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
	return r.name
}

func (r *role) GetNameWithTroop() string {
	return fmt.Sprintf("[%d]%s", r.troopID, r.name)
}

func (r *role) GetStatus() string {
	return fmt.Sprintf(
		"HP: %d, MP: %d, STR: %d, State: %s",
		r.hp, r.mp, r.str, r.state.GetName(),
	)
}

func (r *role) GetAllSkillName() string {
	var list []string

	for i, skill := range r.skills {
		list = append(list, fmt.Sprintf("(%d) %s", i, skill.GetName()))
	}

	return strings.Join(list, " ")
}

func (r *role) AddRelation(ro RelationObserver) bool {
	for _, relation := range r.relations {
		if relation.GetName() == ro.GetName() {
			return false
		}
	}

	r.relations = append(r.relations, ro)
	return true
}

func (r *role) GetRelations() []RelationObserver {
	return r.relations
}
