package action

import (
	"fmt"
	"strconv"

	"rpg/internal/domain"
)

type hero struct{}

func NewHero() domain.ActionStrategy {
	return &hero{}
}

func (h hero) S1(skillsIDs []int) int {
	var n int
	var input string

	for {
		_, _ = fmt.Scanln(&input)
		n, _ = strconv.Atoi(input)

		if n >= len(skillsIDs) {
			fmt.Println("選擇超出範圍，請重新選擇。")
			continue
		} else {
			break
		}
	}

	return skillsIDs[n]
}

func (h hero) S2(availableIDs []int, limit int) []int {
	if len(availableIDs) == limit {
		return availableIDs
	}

	var targets []int
	var input string

	for len(targets) != limit {
		_, _ = fmt.Scanln(&input)
		n, _ := strconv.Atoi(input)

		if n >= len(availableIDs) {
			fmt.Println("選擇超出範圍，請重新選擇。")
			continue
		} else {
			targets = append(targets, availableIDs[n])
		}
	}

	return targets
}
