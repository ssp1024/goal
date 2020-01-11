package goal_test

import (
	"fmt"
	"time"

	"github.com/sweetycode/goal"
)

func ExampleTime() {
	cost := goal.Time(func() {
		time.Sleep(3 * time.Second)
	})

	fmt.Println(int(cost / time.Second))
	// Output:
	// 3
}
