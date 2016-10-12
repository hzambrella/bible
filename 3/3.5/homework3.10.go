// 练习 3.10： 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。
// 比如12345  变为 12，345
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for x, y := range s {
		if n%3 == 0 && x%3 == 0 && x != 0 {
			buf.WriteString(",")
		}
		if n%3 == 1 && x%3 == 1 && x != 0 {
			buf.WriteString(",")
		}
		if n%3 == 2 && x%3 == 2 && x != 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%c", y)
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("dasfadsfadsfaddgadsgaa"))
}
