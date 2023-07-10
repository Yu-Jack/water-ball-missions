package action

import (
	"fmt"
	"strconv"
	"strings"

	"rpg/internal/domain"
)

type heroTxt struct {
	action
	input []string
}

func NewHeroTxt(input []string) domain.ActionStrategy {
	return &heroTxt{
		input: input,
		action: action{
			name: domain.StrategyNameHero,
		},
	}
}

func (h *heroTxt) popInput() string {
	first := h.input[0]
	h.input = h.input[1:]
	return first
}

func (h *heroTxt) S1(skillsIDs []int) int {
	var n int

	for {
		n, _ = strconv.Atoi(h.popInput())

		if n >= len(skillsIDs) {
			fmt.Println("選擇超出範圍，請重新選擇。")
			continue
		} else {
			break
		}
	}

	return skillsIDs[n]
}

func (h *heroTxt) S2(availableIDs []int, limit int, list string) []int {
	if len(availableIDs) == limit {
		return availableIDs
	}

	fmt.Printf(
		"選擇 %d 位目標: %s\n", limit, list,
	)

	var (
		input  string
		inputs []string
		orders []int
		again  = true
	)

	for again {
		for input == "" {
			input = strings.TrimSuffix(h.popInput(), "\n")
			inputs = strings.Split(input, ", ")
		}

		for _, input := range inputs {
			integer, _ := strconv.Atoi(input)
			if integer >= len(availableIDs) {
				fmt.Println("選擇超出範圍，請重新選擇。")
				again = true
				break
			}
			orders = append(orders, integer)
			again = false
		}
	}

	return orders
}
