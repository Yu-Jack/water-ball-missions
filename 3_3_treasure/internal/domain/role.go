//go:generate go-enum -f=$GOFILE --nocase

package domain

import (
	"fmt"
)

const (
	fullHP = 300
)

type Role interface {
	SetState(state State)
	MinusHP(hp int)
	PlusHP(hp int)
	GetHP() int
	SetActionCount(int)
	SetActionOptions([]ActionOption)
	SetPower(power int)
	SetDirections(directions []Direction)
	SetAttackScope(scope AttackScope)
	TakeTurn()
	TakeTreasure(role Role, treasure Treasure)
	GetPosition() Position
	Removed()
	Damage(attacked Role)
	DefaultScope() AttackScope
	SetMap(m *Map)
	GetName() string
	SetPosition(position Position)
	RandomMoved()
	AttackedTransfer()
	IsAttackedAble() bool
	IsFullHP() bool
}

// RoleName
/*
ENUM(
Monster
Character
)
*/
type RoleName int

// Direction
/*
ENUM(
Up = "↑"
Down = "↓"
Left = "←"
Right = "→"
)
*/
type Direction string

// AttackScope
/*
ENUM(
One = "單點攻擊"
Line = "一條線攻擊"
All = "全範圍攻擊"
)
*/
type AttackScope string

// ActionOption
/*
ENUM(
Move = "移動"
Attack = "攻擊"
)
*/
type ActionOption string

type role struct {
	m              *Map
	name           RoleName
	state          State
	hp             int
	actionCount    int
	actionOptions  []ActionOption
	directions     []Direction
	power          int
	scope          AttackScope
	position       Position
	controlledAble bool
}

func (r *role) takeTurn(action func()) {
	// 回合開始前
	r.state.Do()
	r.state.CountDown()

	if r.hp <= 0 {
		return
	}
	// 回合開始，執行動作
	ac := r.actionCount
	for i := 0; i < ac; i++ {
		action()
	}

	// 回合結束後確認新狀態是否有變化
	if !r.state.IsNormal() {
		r.m.AppendLog(fmt.Sprintf("角色: %s 的狀態 %s 剩餘 %d\n", r.name, r.state.GetType(), r.state.GetRound()))
	}
	r.state.Recover()
}

func (r *role) GetName() string {
	if r.name == RoleNameMonster {
		return fmt.Sprintf("%s-%d", r.name, monsterNum)
	} else {
		return r.name.String()
	}
}

func (r *role) AttackedTransfer() {
	r.state.AttackedTransfer()
}

func (r *role) RandomMoved() {
	newPosition := r.m.RandomEmptyPosition()
	r.m.DirectlyMove(r.position, newPosition)
	r.position = newPosition
}

func (r *role) SetState(state State) {
	if state.GetType() != r.state.GetType() {
		r.m.AppendLog(fmt.Sprintf("%s 狀態從 %s 變成 %s\n", r.GetName(), r.state.GetType(), state.GetType()))
		r.state.Exit()
	}
	r.state = state
	r.state.Enter() // 觸發進入新狀態的動作
}

func (r *role) IsAttackedAble() bool {
	return r.state.IsAttackedAble()
}

func (r *role) SetActionCount(count int) {
	r.actionCount = count
}

func (r *role) SetActionOptions(options []ActionOption) {
	r.actionOptions = options
}

func (r *role) SetPosition(position Position) {
	r.position = position
}

func (r *role) MinusHP(hp int) {
	beforeHp := r.hp

	r.hp -= hp

	if r.hp <= 0 {
		r.hp = 0
	}

	afterHp := r.hp

	r.m.AppendLog(fmt.Sprintf("%s 血量從 %d 降低到 %d\n",
		r.GetName(), beforeHp, afterHp))

	if r.hp == 0 {
		r.Removed()
	}

}

func (r *role) Damage(attacked Role) {
	if !attacked.IsAttackedAble() {
		r.m.AppendLog(fmt.Sprintf("%s 無敵狀態，不受 %s 攻擊\n", attacked.GetName(), r.GetName()))
		return
	}

	r.m.AppendLog(fmt.Sprintf("%s 使用 %s 攻擊 %s\n", r.GetName(), r.scope, attacked.GetName()))
	attacked.MinusHP(r.power)

	if attacked.GetHP() > 0 {
		attacked.AttackedTransfer()
	}
}

func (r *role) PlusHP(hp int) {
	var afterHp, beforeHp int

	beforeHp = r.hp

	r.hp += hp
	if r.hp >= fullHP {
		r.hp = fullHP
	}

	afterHp = r.hp

	r.m.AppendLog(fmt.Sprintf("%s 血量從 %d 上升到 %d\n",
		r.name, beforeHp, afterHp))

}

func (r *role) GetHP() int {
	return r.hp
}

func (r *role) Removed() {
	r.m.AppendLog(fmt.Sprintf("%s 死亡\n", r.GetName()))
	r.m.remove(r.position)
}

func (r *role) SetDirections(directions []Direction) {
	r.directions = directions
}

func (r *role) GetPosition() Position {
	return r.position
}

func (r *role) SetMap(m *Map) {
	r.m = m
}

func (r *role) SetPower(power int) {
	r.power = power
}

func (r *role) SetAttackScope(scope AttackScope) {
	r.scope = scope
}

func (r *role) TakeTreasure(role Role, treasure Treasure) {
	s := treasure.GiveState()
	s.SetRole(role)
	r.m.AppendLog(fmt.Sprintf("%s 取得 %s\n", role.GetName(), treasure.GetContent()))
	r.SetState(s)
}

func (r *role) IsFullHP() bool {
	return r.hp == fullHP
}
