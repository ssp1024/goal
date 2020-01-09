package goal

import "time"

func Time(fn func()) time.Duration {
	st := time.Now()
	fn()
	return time.Now().Sub(st)
}
