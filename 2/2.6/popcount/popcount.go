package popcount
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
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}


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
