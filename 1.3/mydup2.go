package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Println(err)
			}
			countLines(f, counts)
		}
	}
	for u, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, u)
		}
	}
}
func countLines(f *os.File, c map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		c[input.Text()]++
	}
}
