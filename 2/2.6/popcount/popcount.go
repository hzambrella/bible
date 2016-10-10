package popcount

import "fmt"

var pc [256]byte

// 预生成辅助表格，这是编程中常用的技术
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// 例子代码
// 使用辅助表格,比如pc[255],很快就得到8，即255的二进制有8个1
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// homework2.3
// 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能
func PopCountCircle(x uint64) int {
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

func PopPrint(x uint64) {
	for i := 0; i < 8; i++ {
		fmt.Printf("pc[%v]=%v\n", byte(x>>uint(i*8)), pc[byte(x>>uint(i*8))])
	}
}

// homework2.4
//  用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。
func PopCountMove(n uint64) int {
	var c uint64
	// n>>=1,"="!!!!!
	for c = 0; n > 0; n >>= 1 {
		c += n & 1
	}
	return int(c)
}

// homework2.5
// 表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。
// 为什么n &= (n – 1)能清除最右边的1呢？因为从二进制的角度讲，n相当于在n - 1的最低位加上1。举个例子，8（1000）= 7（0111）+ 1（0001），所以8 & 7 = （1000）&（0111）= 0（0000），清除了8最右边的1（其实就是最高位的1，因为8的二进制中只有一个1）。再比如7（0111）= 6（0110）+ 1（0001），所以7 & 6 = （0111）&（0110）= 6（0110），清除了7的二进制表示中最右边的1（也就是最低位的1）
func PopCountFast(n uint64) int {
	var c uint64
	for c = 0; n > 0; c++ {
		//	n & (n - 1)
		n = n & (n - 1)
	}
	return int(c)
}
