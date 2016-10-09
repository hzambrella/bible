// 对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出。我们会展示一个名为dup的程序的三个版本；灵感来自于Unix的uniq命令，其寻找相邻的重复行。该程序使用的结构和包是个参考范例，可以方便地修改。
//dup的第一个版本打印标准输入中多次出现的行，以重复次数开头。该程序将引入if语句，map数据类型以及bufio包。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // count
	// Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词,该变量从程序的标准输入中读取内容。每次调用input.Scanner，即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到。Scan函数在读到一行时返回true，在无输入时返回false。

	input := bufio.NewScanner(os.Stdin) //os.Stdin,就是在标准输入(终端输入),NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。ScanLines是用于Scanner类型的分割函数（符合SplitFunc）， 本函数会将每一行文本去掉末尾的换行标记作为一个token返回。 返回的行可以是空字符串。换行标记为一个可选的回车后跟一个必选的换行符。 最后一行即使没有换行符也会作为一个token返回。
	// _____________________________________
	// after input,use"ctrl+D"!!!!!!!!
	// _____________________________________
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) //\t 制表符
		}
	}
}

/*Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。

成功调用的Scan方法会逐步提供文件的token， 跳过token之间的字节。token由SplitFunc类型的分割函数指定； 默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。 本包预定义的分割函数可以将文件分割为行、字节、unicode码值、空白分隔的word。 调用者可以定制自己的分割函数。

扫描会在抵达输入流结尾、遇到的第一个I/O错误、token过大不能保存进缓冲时， 不可恢复的停止。当扫描停止后，当前读取位置可能会远在最后一个获得的token后面。 需要更多对错误管理的控制或token很大，或必须从reader连续扫描的程序， 应使用bufio.Reader代替*/
