package goal

import (
	"bytes"
	"fmt"
	"time"
)

type Stat interface {
	String() string
	Incr(string) Stat
	IncrN(string, int) Stat
	Stop()
}

type statItem struct {
	key string
	val int
}

type StatVal struct {
	n int
}

func (v *StatVal) Value() int {
	return v.n
}

func (v *StatVal) String() string {
	return Itoa(v.n)
}

type StatKV map[string]*StatVal

func (kv StatKV) Serialize() string {
	buffer := &bytes.Buffer{}
	buffer.WriteString("[")

	if kv != nil {
		for k, v := range kv {
			buffer.WriteString(k)
			buffer.WriteString(":")
			buffer.WriteString(v.String())
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}

func (v *StatVal) Incr(n int) {
	v.n += n
}

func NewStat() Stat {
	return newStat(0, nil)
}

func NewLogStat(duration time.Duration) Stat {
	return newStat(duration, cached(func(s string) {
		Infof(s)
	}))
}

func NewPrintStat(duration time.Duration) Stat {
	return newStat(duration, cached(func(s string) {
		fmt.Println(s)
	}))
}

func cached(fn func(string)) func(StatKV, bool) {
	var cache string
	return func(kv StatKV, dirty bool) {
		if dirty {
			cache = kv.Serialize()
		}
		fn(cache)
	}
}

func newStat(duration time.Duration, fn func(StatKV, bool)) Stat {
	stat := &statImpl{
		duration: duration,
		fn:       fn,

		kv:       make(StatKV),
		shadowKv: nil,
		ch:       make(chan statItem, 10000),
		waitStop: NewWaitStop(),
	}
	go stat.collectorThread()
	return stat
}

type statImpl struct {
	duration time.Duration
	fn       func(kv StatKV, dirty bool)

	kv       StatKV
	shadowKv map[string]int
	ch       chan statItem

	waitStop *WaitStop
}

func (stat *statImpl) collectorThread() {
	defer stat.waitStop.Stop()

	var tickerCh <-chan time.Time
	if stat.duration > 0 && stat.fn != nil {
		ticker := time.NewTicker(stat.duration)
		defer ticker.Stop()
		tickerCh = ticker.C
	}

	dirty := false

	for {
		select {
		case item, ok := <-stat.ch:
			if !ok {
				return
			}
			dirty = true
			val, ok := stat.kv[item.key]
			if !ok {
				val = &StatVal{}
				stat.kv[item.key] = val
			}
			val.Incr(item.val)
		case <-tickerCh:
			stat.fn(stat.kv, dirty)
			dirty = false
		}
	}
}

func (stat *statImpl) IncrN(key string, n int) Stat {
	stat.ch <- statItem{
		key: key,
		val: n,
	}

	return stat
}

func (stat *statImpl) Incr(key string) Stat {
	return stat.IncrN(key, 1)
}

func (stat *statImpl) Stop() {
	close(stat.ch)

	stat.waitStop.Wait()
}

func (stat *statImpl) String() string {
	buffer := &bytes.Buffer{}
	buffer.WriteString("Stat")

	if stat.waitStop.Stopped() {
		buffer.WriteString(stat.kv.Serialize())
	} else {
		buffer.WriteString("[running...]")
	}
	return buffer.String()
}
