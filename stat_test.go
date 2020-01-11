package goal_test

import (
	"fmt"
	"time"

	"github.com/sweetycode/goal"
)

func ExampleStat() {
	stat := goal.NewPrintStat(time.Millisecond * 200)
	defer stat.Stop()

	go func() {
		stat.Incr("counter")
		time.Sleep(time.Millisecond * 300)
		stat.Incr("counter")
	}()

	time.Sleep(500 * time.Millisecond)
	//Output:
	//Stat[counter:1, ]
	//Stat[counter:2, ]
}

func ExampleStat_stop() {
	stat := goal.NewPrintStat(time.Millisecond * 200)

	stat.Incr("counter")
	stat.Incr("counter")

	fmt.Println(stat)
	stat.Stop()
	fmt.Println(stat)
	//Output:
	// Stat[running...]
	// Stat[counter:2, ]
}
