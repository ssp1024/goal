package goal_test

import (
	"fmt"

	"github.com/sweetycode/goal"
)

func ExampleMaxInt() {
	n := goal.MaxInt(10, 100, 50)
	fmt.Println(n)
	//Output:
	// 100
}
