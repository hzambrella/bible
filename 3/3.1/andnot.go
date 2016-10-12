package main
// &^     位清空 (AND NOT)

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d&^%d=%d\n", i, j, i&^j)
		}
	}
}
