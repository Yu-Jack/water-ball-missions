package action

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	reader := bufio.NewReader(os.Stdin)

	var (
		input  string
		inputs []string
		orders []int
		again  = true
	)

	for again {
		for input == "" {
			input, _ = reader.ReadString('\n')
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
