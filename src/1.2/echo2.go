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
