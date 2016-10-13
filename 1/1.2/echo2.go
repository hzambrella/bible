// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。

package main
import(
	"os"
	"fmt"
)
func main(){
	s:=os.Args[0:]
	for k,v:=range s {
		fmt.Println(k,v)
	}
}
