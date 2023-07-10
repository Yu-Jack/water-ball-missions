package action

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"rpg/internal/domain"
)

type hero struct {
	reader *bufio.Reader
}

func NewHero(reader *bufio.Reader) domain.ActionStrategy {
	return &hero{reader: reader}
}

func (h *hero) S1(skillsIDs []int) int {
	var n int
	var input string

	for {
		input, _ = h.reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
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

func (h *hero) S2(availableIDs []int, limit int, list string) []int {
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
			input, _ = h.reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")
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
