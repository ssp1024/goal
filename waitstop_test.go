package goal_test

import (
	"fmt"
	"time"

	"github.com/sweetycode/goal"
)

func ExampleWaitStop() {
	waitStop := goal.NewWaitStop()

	go func() {
		defer waitStop.Stop()

		fmt.Println("start async process")
		time.Sleep(1 * time.Second)
		fmt.Println("end async process")
	}()

	if waitStop.Stopped() {
		fmt.Println("async stopped")
	}

	waitStop.Wait()
	fmt.Println("wait finished")

	// Output:
	// start async process
	// end async process
	// wait finished
}
