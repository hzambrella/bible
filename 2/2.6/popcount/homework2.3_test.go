package popcount

import "testing"
import "fmt"


// homework2.4
func TestPopCountMove(t *testing.T) {
	result := PopCountMove(66666)
	fmt.Println(result)
}
// homework2.5
func TestPopCountFast(t *testing.T) {
	result := PopCountFast(254)
	fmt.Println(result)
}
