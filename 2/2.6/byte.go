// pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数，这样的话在处理64bit宽度的数字时就没有必要循环64次，只需要8次查表就可以了。（这并不是最快的统计1bit数目的算法，但是它可以方便演示init函数的用法，并且演示了如果预生成辅助表格，这是编程中常用的技术）。
// 此程序用来统计数字的二进制形式中有几个1
// 用于处理每个8bit宽度的数字含二进制的1bit的bit个数.
package main

import (
	"fmt"
)

func main() {
	researchmove()
	move(255)
	// table()
}

// 查表法统计二进制数的1的个数当中，研究表
func table() {
	//  byte array of 256 length
	var pc [256]byte
	fmt.Println("pc is ", pc)
	fmt.Println("pc[i] = pc[i/2] + byte(i&1)")
	for i := range pc {
		// what is  it?
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("%d:%v\t", i, byte(i))
		fmt.Printf("byte(%d&1)=%v\t", i, byte(i&1))
		fmt.Printf("pc[%d]:=%v\n", i, pc[i])
	}
}

// 普通法统计二进制数的1的个数
func move(n uint64) int {
	var c uint64
	//	for c = 0; n > 0; n >> 1 {
	// ./byte.go:30: n >> 1 evaluated but not used
	for c = 0; n > 0; n >>= 1 {
		c += n & 1
		fmt.Printf("n=%v\n", n)
	}
	return int(c)
}

// n>>1和n>>=1
func researchmove() {
	var n int = 255
	fmt.Println(n >> 1)
	fmt.Println(n)
	n >>= 1
	fmt.Println(n)
}
