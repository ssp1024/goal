package main

import (
	"fmt"
	"time"

	"github.com/sweetycode/goal"
)

func main() {
	waitStop := goal.NewWaitStop()

	stat := goal.NewPrintStat(time.Second * 1)

	go func() {
		for i := 0; i < 1000; i++ {
			stat.Incr("test")
			time.Sleep(5 * time.Millisecond)
		}
		waitStop.Stop()
	}()

	fmt.Println(stat)

	waitStop.Wait()

	stat.Stop()
	fmt.Println(stat)
}
