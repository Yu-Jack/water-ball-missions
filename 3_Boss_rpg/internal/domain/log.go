package domain

import "fmt"

func LogDamage(current, target string, damage int) {
	fmt.Printf(
		"%s 對 %s 造成 %d 點傷害。\n",
		current, target, damage,
	)
}
