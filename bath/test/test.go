package main

import (
	"fmt"

	"github.com/davyxu/cellnet/bath/comm"
)

func main() {
	// num := uint64(20)
	// "sau", "spy", "sau1spy", "sau0spy"
	var n1, n2, n3, n4 int
	for i := 0; i < 1; i++ {
		item := comm.Prob()
		switch item {
		case "sau":
			n1++
		case "spy":
			n2++
		case "sau1spy":
			n3++
		case "sau0spy":
			n4++
		}

	}
	fmt.Println(n1, n2, n3, n4)

}
