package main

import (
	"fmt"

	"relationship/internal/domain"
)

func Example() {
	analyzer := domain.NewRelationShipAnalyzerAdapter()
	graph := analyzer.Parse(`A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z`)

	fmt.Println(analyzer.GetMutualFriends("A", "B"))
	fmt.Println(graph.HasConnection("A", "F"))
	fmt.Println(graph.HasConnection("A", "D"))

	// Output:
	// [D]
	// false
	// true
}
