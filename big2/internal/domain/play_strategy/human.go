package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"big2/internal/domain"
)

type human struct{}

func NewHumanStrategy() domain.InputStrategy {
	return &human{}
}

func (h human) Input() ([]int, bool) {
	reader := bufio.NewReader(os.Stdin)
	var input string
	var inputs []string

	for input == "" {
		fmt.Print("請輸入牌號：")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSuffix(input, "\n")
		inputs = strings.Split(input, " ")
	}

	if inputs[0] == "-1" {
		return []int{}, true
	}

	orders := make([]int, 0, len(inputs))
	for _, input := range inputs {
		integer, _ := strconv.Atoi(input)
		orders = append(orders, integer)
	}

	return orders, false
}

func (h human) Name() string {
	var name string

	fmt.Println("請輸入你的名字：")
	_, _ = fmt.Scan(&name)
	fmt.Print("你好, " + name + "!\n")

	return name
}
