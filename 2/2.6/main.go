package main

import "fmt"
import "bible/2/2.6/popcount"

func main() {
	y := uint64(123131255)
	// byte类型只有8bit，高出的位数会被省略,移位可以看到高位的,不信看下面的

	fmt.Println(byte(y))
	fmt.Println(byte(y >> 8))
	x := popcount.PopCount(y)
	fmt.Println(x)
	popcount.PopPrint(y)
}
