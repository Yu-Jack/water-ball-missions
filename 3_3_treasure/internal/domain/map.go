//go:generate go-enum -f=$GOFILE --nocase

package domain

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Map struct {
	m                 [][]MapObject
	xRayLimit         int
	yRayLimit         int
	character         Role
	monsters          []Role
	monsterGenerator  func()
	treasures         []Treasure
	treasureGenerator func()
	logs              []string
}

type Position struct {
	XRay, YRay int
}

type Obstacle struct {
	Position Position
}

func (o *Obstacle) GetPosition() Position {
	return o.Position
}

// MapObject
/*
ENUM(
Obstacle = "O"
Monster = "M"
Character = "C"
treasure = "x"
Empty = " "
Wall = "W"
)
*/
type MapObject string

func NewMap() *Map {
	xRayLimit, yRayLimit := 10, 10

	m := make([][]MapObject, yRayLimit)
	for i := 0; i < yRayLimit; i++ {
		m[i] = make([]MapObject, xRayLimit)
	}

	return &Map{
		m:         m,
		xRayLimit: xRayLimit,
		yRayLimit: yRayLimit,
		monsters:  []Role{},
		treasures: []Treasure{},
		logs:      []string{},
	}
}

func (m *Map) SetRole(r Role) {
	r.SetMap(m)

	p := r.GetPosition()

	if strings.Contains(r.GetName(), RoleNameMonster.String()) {
		m.setMapLabel(p, MapObjectMonster)
		m.monsters = append(m.monsters, r)
	}

	if r.GetName() == RoleNameCharacter.String() {
		m.setMapLabel(p, MapObjectCharacter)
		m.character = r
	}
}

func (m *Map) SetUpMonsterGenerator(generator func()) {
	m.monsterGenerator = generator
}

func (m *Map) SetUpTreasureGenerator(generator func()) {
	m.treasureGenerator = generator
}

func (m *Map) GenerateMonster() {
	m.monsterGenerator()
}

func (m *Map) GenerateTreasure() {
	m.treasureGenerator()
}

func (m *Map) SetTreasure(t Treasure) {
	m.treasures = append(m.treasures, t)
	m.setMapLabel(t.GetPosition(), MapObjectTreasure)
}

func (m *Map) SetObstacle(o Obstacle) {
	m.setMapLabel(o.GetPosition(), MapObjectObstacle)
}

func (m *Map) setMapLabel(p Position, object MapObject) {
	m.m[p.YRay][p.XRay] = object
}

func (m *Map) AppendLog(str string) {
	s := fmt.Sprintf("系統時間: %s, 系統訊息 %s", time.Now().Format(time.RFC1123), str)
	m.logs = append([]string{s}, m.logs...)
}

func (m *Map) Start() {
	defer func() {
		fmt.Println("遊戲結束")
		m.Print()
	}()

	for m.character != nil && len(m.monsters) != 0 {
		m.GenerateTreasure()
		m.GenerateMonster()

		m.character.TakeTurn()

		for _, monster := range m.monsters {
			monster.TakeTurn()

			if m.character == nil {
				return
			}
		}
	}
}

func (m *Map) DirectlyMove(oldPosition, newPosition Position) {
	m.changePosition(m.character, oldPosition, newPosition, Direction(m.m[oldPosition.YRay][oldPosition.XRay]))
}

func (m *Map) RandomEmptyPosition() Position {
	isEmpty := false
	var x, y int

	for !isEmpty {
		y, x = rand.Int()%m.yRayLimit, rand.Int()%m.xRayLimit
		if m.m[y][x] == "" {
			isEmpty = true
			break
		}
	}

	return Position{XRay: x, YRay: y}
}

func (m *Map) Move(role Role, position Position, direction Direction) (successMove bool) {
	originalP := position

	switch direction {
	case DirectionUp:
		if walkAble, _ := m.isWalkAble(position.XRay, position.YRay-1); !walkAble {
			return false
		}
		position.YRay -= 1
	case DirectionDown:
		if walkAble, _ := m.isWalkAble(position.XRay, position.YRay+1); !walkAble {
			return false
		}
		position.YRay += 1
	case DirectionLeft:
		if walkAble, _ := m.isWalkAble(position.XRay-1, position.YRay); !walkAble {
			return false
		}
		position.XRay -= 1
	case DirectionRight:
		if walkAble, _ := m.isWalkAble(position.XRay+1, position.YRay); !walkAble {
			return false
		}
		position.XRay += 1
	}

	// success move
	if m.m[position.YRay][position.XRay] == "" {
		m.changePosition(role, originalP, position, direction)
		role.SetPosition(position)
		return true
	}

	// trigger touch treasure
	if m.m[position.YRay][position.XRay] == MapObjectTreasure {
		role.TakeTreasure(role, m.getTreasure(position))

		// 移除物件
		m.remove(position)

		// 移動過去
		m.changePosition(role, originalP, position, direction)
		role.SetPosition(position)
		return true
	}

	return true
}

func (m *Map) changePosition(role Role, originalP Position, position Position, direction Direction) {
	m.m[position.YRay][position.XRay] = m.m[originalP.YRay][originalP.XRay]
	m.m[originalP.YRay][originalP.XRay] = ""

	if role.GetName() == RoleNameCharacter.String() {
		m.m[position.YRay][position.XRay] = MapObject(direction.String())
	}
}

func (m *Map) getTreasure(position Position) Treasure {
	for _, t := range m.treasures {
		p := t.GetPosition()
		if p.YRay == position.YRay && p.XRay == position.XRay {
			return t
		}
	}
	return nil
}

func (m *Map) isCharacterAround(position Position) (bool, Role) {
	// 跟主要角色距離為 1
	p := m.character.GetPosition()
	isAround := (abs(position.YRay-p.YRay) + abs(position.XRay-p.XRay)) == 1

	return isAround, m.character
}

func (m *Map) GetMonster(position Position) Role {
	for _, r := range m.monsters {
		if r.GetPosition().YRay == position.YRay && r.GetPosition().XRay == position.XRay {
			return r
		}
	}
	return nil
}

func (m *Map) GetRolesBasedOneDirection(position Position, direction Direction) []Role {
	var roles []Role
	var start, end int
	var checkCondition func(i int) bool
	var checkOneByBone func(i int) int

	if direction == DirectionUp || direction == DirectionDown {

		if direction == DirectionUp {
			start = position.YRay - 1
			end = 0

			checkCondition = func(i int) bool {
				return i >= end
			}
			checkOneByBone = func(i int) int {
				return i - 1
			}
		} else {
			start = position.YRay + 1
			end = m.yRayLimit - 1

			checkCondition = func(i int) bool {
				return i <= end
			}
			checkOneByBone = func(i int) int {
				return i + 1
			}
		}

		for i := start; checkCondition(i); i = checkOneByBone(i) {
			if m.m[i][position.XRay] == MapObjectObstacle {
				break
			}

			if m.m[i][position.XRay] == MapObjectMonster {
				roles = append(roles, m.GetMonster(Position{position.XRay, i}))
			}
		}
	}

	if direction == DirectionLeft || direction == DirectionRight {
		if direction == DirectionLeft {
			start = position.XRay - 1
			end = 0

			checkCondition = func(i int) bool {
				return i >= end
			}
			checkOneByBone = func(i int) int {
				return i - 1
			}
		} else {
			checkCondition = func(i int) bool {
				return i <= end
			}
			checkOneByBone = func(i int) int {
				return i + 1
			}
			start = position.XRay + 1
			end = m.xRayLimit - 1
		}

		for i := start; checkCondition(i); i = checkOneByBone(i) {
			if m.m[position.YRay][i] == MapObjectObstacle {
				break
			}

			if m.m[position.YRay][i] == MapObjectMonster {
				roles = append(roles, m.GetMonster(Position{i, position.YRay}))
			}
		}
	}

	return roles
}

func (m *Map) GetAllRolesExcludePosition(excludePosition Position) []Role {
	var roles []Role

	for _, r := range m.monsters {
		if r.GetPosition().YRay == excludePosition.YRay && r.GetPosition().XRay == excludePosition.XRay {
			continue
		}
		roles = append(roles, r)
	}

	if m.character != nil &&
		!(m.character.GetPosition().YRay == excludePosition.YRay &&
			m.character.GetPosition().XRay == excludePosition.XRay) {
		roles = append(roles, m.character)
	}

	return roles

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m *Map) Print() {
	for i := 0; i < m.yRayLimit; i++ {
		row := ""
		for j := 0; j < m.xRayLimit; j++ {
			s := string(m.m[i][j])
			if s == "" {
				row += "-"
			} else {
				row += s
			}
		}
		if len(m.logs)-1 >= i {
			fmt.Printf("%s        %s", row, m.logs[i])
		} else {
			fmt.Printf("%s\n", row)
		}
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++")
}

func (m *Map) remove(position Position) {
	defer func() {
		m.m[position.YRay][position.XRay] = ""
	}()

	for i, t := range m.treasures {
		p := t.GetPosition()
		if p.YRay == position.YRay && p.XRay == position.XRay {
			m.treasures = append(m.treasures[:i], m.treasures[i+1:]...)
			return
		}
	}

	for i, t := range m.monsters {
		p := t.GetPosition()
		if p.YRay == position.YRay && p.XRay == position.XRay {
			m.monsters = append(m.monsters[:i], m.monsters[i+1:]...)
			return
		}
	}

	if p := m.character.GetPosition(); p.YRay == position.YRay && p.XRay == position.XRay {
		m.character = nil
	}
}

func (m *Map) isWalkAble(xRay, yRay int) (bool, MapObject) {
	if xRay < 0 || xRay >= m.xRayLimit {
		return false, MapObjectWall
	}

	if yRay < 0 || yRay >= m.yRayLimit {
		return false, MapObjectWall
	}

	if m.m[yRay][xRay] == MapObjectObstacle {
		return false, MapObjectObstacle
	}

	if m.m[yRay][xRay] == MapObjectMonster {
		return false, MapObjectMonster
	}

	if m.m[yRay][xRay] == MapObjectCharacter {
		return false, MapObjectCharacter
	}

	return true, MapObjectEmpty
}
