// 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）
// 写一个单元测试,运行go test -bench=.
package main

import "fmt"

var pc [256]byte

// 预生成辅助表格，这是编程中常用的技术
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 使用辅助表格,比如pc[255],很快就得到8，即255的二进制有8个1
func PopCount(x uint64) int {
	var sum int
	for i := 0; i < 8; i++ {
		k := i * 8
		//	sum=sum+pc[byte(x>>k)]
		// ./homework2.3.go:18: invalid operation: x >> k (shift count type int, must be unsigned integer)
		//	sum=sum+pc[byte(x>>uint(k))]
		// ./homework2.3.go:20: invalid operation: sum + pc[byte(x >> uint(k))] (mismatched types int and byte)
		// 这里注意！！！！！！
		// pc[i] is byte, so change to int
		sum += int(pc[byte(x>>uint(k))])
	}
	return sum
}

func main() {
	y := uint64(123131255)
	// byte类型只有8bit，高出的位数会被省略,移位可以看到高位的,不信看下面的

	fmt.Println(byte(y))
	fmt.Println(byte(y >> 8))
	x := PopCount(y)
	fmt.Println(x)
	PopPrint(y)
}
func PopPrint(x uint64) {
	for i := 0; i < 8; i++ {
		fmt.Printf("pc[%v]=%v\n", byte(x>>uint(i*8)), pc[byte(x>>uint(i*8))])
	}
}
