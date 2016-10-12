// 练习 3.12： 编写一个函数，判断两个字符串是否是是相互打乱的，也就是说它们有着相同的字符，但是对应不同的顺序。
package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	if upset("abc", "cab") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func upset(str1, str2 string) bool {
	var yes bool = true
	a1 := statistic(str1)
	a2 := statistic(str2)
	if len(a1) != len(a2) {
		yes = false
		return yes
	}
	for k, v := range a1 {
		if v != a2[k] {
			yes = false
			break
		}
	}
	return yes
}

func statistic(str string) map[byte]int {
	// var num map[byte]int
	num:=make(map[byte]int)
	var x byte
	for i := 0; i < len(str); i++ {
		x = str[i]
		num = add(num, x)
	}
	return num
}

func add(num map[byte]int, x byte) map[byte]int {

	defer mu.Unlock()
	mu.Lock()
	num[x]++
	return num
}
