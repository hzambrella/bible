// 函数的功能是将一个表示整值的字符串，每隔三个字符插入一个逗号分隔符，例如“12345”处理后成为“12,345”。这个版本只适用于整数类型；支持浮点数类型的支持留作练习。
package main

import (
	"fmt"
)

func comma(s string) string {
	n := len(s)
	//	if n<3 {
	if n <= 3 {
		return s
	}
	// 递归
	return comma(s[:n-3]) + "," + s[n-3:n]
}

func main() {
	fmt.Println(comma("dasfadsfadsfaddgadsg"))
}
