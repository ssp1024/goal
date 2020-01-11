package goal

import "time"

//Time name after linux bash `time` command, return duration of `fn` execution
func Time(fn func()) time.Duration {
	st := time.Now()
	fn()
	return time.Now().Sub(st)
}
