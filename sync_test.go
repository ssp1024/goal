package goal_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/sweetycode/goal"
)

func ExampleSelectWaitGroup() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("start async process")
		time.Sleep(3 * time.Second)
		//unable reach
	}()

	select {
	case <-goal.SelectWaitGroup(wg):
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	// Output:
	// start async process
	// timeout
}
