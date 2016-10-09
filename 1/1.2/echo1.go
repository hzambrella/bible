// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。

// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main
import(
	"fmt"
	"os"
)
func main(){
	var s string
	for i:=0;i<len(os.Args);i++{
		s+=os.Args[i]
	}
	fmt.Println(s)
}
