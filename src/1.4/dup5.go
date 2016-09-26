// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。
// // map[string]*struct
package main

import (
	"bufio"
	"fmt"
	"os"
)

type content struct {
	num      int
	filename []string
}

func main() {
	counts := make(map[string]*content)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			fname := f.Name()
			countLines(f, counts, fname)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.num > 1 {
			fmt.Printf("%d\t%s\n", n.num, line)
	}
}

func countLines(f *os.File, counts map[string]*content, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()].num++
		for _, v := range counts[input.Text()].filename {

			if v != fileName {
				counts[input.Text()].filename = append(counts[input.Text()].filename, v)
			}
		}
	}
}
