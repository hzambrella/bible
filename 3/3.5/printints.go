package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		fmt.Println(i)
		//	if i <len(values) {      //"[,1,2,3]"
		if i > 0 { //values[0],values[1],values[2]
			buf.WriteByte(',')
			// "," is string
			//		buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
