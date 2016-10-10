// flag and pointer
package main

import (
	"flag"
	"fmt"
	"strings"
)

// define a tag.输入 -n -s 有惊喜
var n = flag.Bool("n", false, "omit trailing newline")
var s = flag.String("s", " ", "separator")

func main() {
	// 来解析命令行参数并传入到定义好的标签。
	flag.Parse()
	// fmt.Println(strings.Join(flag.Args(),*s))    // 有换行符
	// 在解析之后，标签对应的参数可以从flag.Args()获取到
	// 一个标签的解析会在下次出现第一个非标签参数（“-”就是一个非标签参数）的时候停止，或者是在终止符号“--”的时候停止
	fmt.Print(strings.Join(flag.Args(), *s))

	if !*n {
		fmt.Println()
	}
	// 输入-n时，n为true
	nn, err := fmt.Println(*n)
	fmt.Println(nn)
	if err != nil {
		fmt.Println(err)
	}
}
