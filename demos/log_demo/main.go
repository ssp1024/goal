package main

import "github.com/sweetycode/goal"

func main() {
	goal.Debugf("NOTHING")
	goal.Infof("info msg")
	goal.Warnf("warn msg")

	goal.SetLogLevel(goal.DebugLevel)

	goal.Debugf("debug msg")
}
