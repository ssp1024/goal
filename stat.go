package goal

import (
	"bytes"
	"fmt"
	"time"
)

//Stat provide efficient way to collect application statistics. can used in high concurrency situation.
//
//Internal implemention avoid use of synchronize operation,  the tride off is cannot retrive metrics value.
type Stat interface {
	//Get stringify of statistics.
	String() string
	//Incr statistics `key` by 1
	Incr(key string) Stat
	//Incr statistics `key` by n
	IncrN(key string, n int) Stat
	//Stop stat
	Stop()
}

type statItem struct {
	key string
	val int
}

type statVal struct {
	n int
}

func (v *statVal) Value() int {
	return v.n
}

func (v *statVal) String() string {
	return Itoa(v.n)
}

func (v *statVal) Incr(n int) {
	v.n += n
}

type statKV map[string]*statVal

func (kv statKV) Serialize() string {
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

//NewStat create new instance.
func NewStat() Stat {
	return newStat(0, nil)
}

//NewLogStat create new instance, and log stat every `duration` interval.
func NewLogStat(duration time.Duration) Stat {
	return newStat(duration, cached(func(s string) {
		Infof(s)
	}))
}

//NewPrintStat create new instance, and print stat every `duration` interval.
func NewPrintStat(duration time.Duration) Stat {
	return newStat(duration, cached(func(s string) {
		fmt.Println(s)
	}))
}

func cached(fn func(string)) func(statKV, bool) {
	var cache string
	return func(kv statKV, dirty bool) {
		if dirty {
			cache = "Stat" + kv.Serialize()
		}
		fn(cache)
	}
}

func newStat(duration time.Duration, fn func(statKV, bool)) Stat {
	stat := &statImpl{
		duration: duration,
		fn:       fn,

		kv:       make(statKV),
		shadowKv: nil,
		ch:       make(chan statItem, 10000),
		waitStop: NewWaitStop(),
	}
	go stat.collectorThread()
	return stat
}

type statImpl struct {
	duration time.Duration
	fn       func(kv statKV, dirty bool)

	kv       statKV
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
				val = &statVal{}
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
