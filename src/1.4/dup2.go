package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// go build dup2.go,then input: " ./dup2 filename ....."
	files := os.Args[1:]
	// if filename is nil , it is dup1.go
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// read file  依次读取文件里面的数据，将其变为os.File型的数据
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// look ! os.File
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
