package domain

import (
	"fmt"
	"strconv"
)

type roleCharacter struct {
	role
}

func NewRoleCharacter(position Position, state State) Role {
	r := &roleCharacter{
		role: role{
			position:      position,
			power:         50,
			scope:         AttackScopeLine,
			hp:            300,
			actionCount:   1,
			actionOptions: []ActionOption{ActionOptionMove, ActionOptionAttack},
			name:          RoleNameCharacter,
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

	return r
}

func (r *roleCharacter) DefaultScope() AttackScope {
	return AttackScopeLine
}

func (r *roleCharacter) TakeTurn() {
	r.m.Print()

	r.takeTurn(func() {
		action := r.chooseAction()

		if action == ActionOptionMove {
			r.actionMove()
		}

		if action == ActionOptionAttack {
			r.actionAttack()
		}
	})
}

func (r *roleCharacter) chooseAction() ActionOption {
	fmt.Println("請選擇行為：")
	for i, s := range r.actionOptions {
		fmt.Printf("%d: %s\n", i, s)
	}
	input := ""
	_, _ = fmt.Scanln(&input)
	n, _ := strconv.Atoi(input)
	action := r.actionOptions[n]
	return action
}

func (r *roleCharacter) actionAttack() {
	var input string
	if r.scope == AttackScopeLine {
		fmt.Println("請選擇攻擊方向：")
		for i, d := range []Direction{
			DirectionUp,
			DirectionDown,
			DirectionLeft,
			DirectionRight,
		} {
			fmt.Printf("%d: %s\n", i, d)
		}
		_, _ = fmt.Scanln(&input)
		n, _ := strconv.Atoi(input)
		direction := r.directions[n%4]

		roles := r.m.GetRolesBasedOneDirection(r.position, direction)
		for _, ro := range roles {
			r.Damage(ro)
		}
	}

	if r.scope == AttackScopeAll {
		fmt.Println("全範圍攻擊")
		roles := r.m.GetAllRolesExcludePosition(r.position)
		for _, ro := range roles {
			r.Damage(ro)
		}
	}
}

func (r *roleCharacter) actionMove() {
	var input string
	fmt.Println("請選擇移動方向：")
	for i, d := range r.directions {
		fmt.Printf("%d: %s\n", i, d)
	}
	_, _ = fmt.Scanln(&input)
	n, _ := strconv.Atoi(input)
	d := r.directions[n%4]
	r.m.Move(r, r.position, d)
}
