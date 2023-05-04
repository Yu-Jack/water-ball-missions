package domain

import (
	"math/rand"
)

var monsterNum = 0

type roleMonster struct {
	role
}

func NewRoleMonster(position Position, state State) Role {
	r := &roleMonster{
		role: role{
			position:      position,
			power:         50,
			scope:         AttackScopeOne,
			hp:            1,
			actionCount:   1,
			actionOptions: []ActionOption{ActionOptionMove, ActionOptionAttack},
			name:          RoleNameMonster,
			directions: []Direction{
				DirectionUp,
				DirectionDown,
				DirectionLeft,
				DirectionRight,
			},
			state: state,
		},
	}

	state.SetRole(r)

	monsterNum++

	return r
}

func (r *roleMonster) DefaultScope() AttackScope {
	return AttackScopeOne
}

func (r *roleMonster) TakeTurn() {
	r.takeTurn(func() {
		if isAround, character := r.m.isCharacterAround(r.position); isAround && !r.state.IsOrderless() {
			if r.scope == AttackScopeOne {
				r.Damage(character)
			}

			if r.scope == AttackScopeAll {
				roles := r.m.GetAllRolesExcludePosition(r.position)
				for _, ro := range roles {
					r.Damage(ro)
				}
			}
		} else {
			d := r.directions[rand.Int()%len(r.directions)] // 隨機一個方向
			r.m.Move(r, r.position, d)
		}
	})
}
