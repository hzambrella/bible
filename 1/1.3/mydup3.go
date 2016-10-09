package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
		// my error:    for_
	// for_,str := range strings.Split(string(data), "\n"){
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	for _, v := range files {
		data, err := ioutil.ReadFile(v)
		if err != nil {
			fmt.Println(err)
			continue // !! I forget it
		}
		str := strings.Split(string(data), "\n")
		for _, v := range str {
			counts[v]++
		}
	}
	for k, v := range counts {
		if v > 1 {
// I should use Printf!!
	//		fmt.Sprintf("%d\t%s", v, k)
	fmt.Printf("%d\t%s\n",v,k)
		}
	}
}
