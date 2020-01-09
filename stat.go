package goal

import (
	"bytes"
	"time"
)

type Stat interface {
	String() string
	Incr(string) Stat
	IncrN(string, int) Stat
}

type statItem struct {
	key string
	val int
}

type statVal int

func NewStat(duration time.Duration) Stat {
	stat := &statImpl{
		duration: duration,
		kv:       make(map[string]*statVal),
		shadowKv: nil,
		ch:       make(chan statItem, 10000),
		waitStop: NewWaitStop(),
	}
	go stat.collectorThread()
	return stat
}

func NewLogStat(duration time.Duration) Stat {
	stat := NewStat(duration)

	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()
		for range ticker.C {
			Infof(stat.String())
		}
	}()

	return stat
}

type statImpl struct {
	duration time.Duration

	kv       map[string]*statVal
	shadowKv map[string]int
	ch       chan statItem

	waitStop *WaitStop
}

func (stat *statImpl) collectorThread() {
	defer stat.waitStop.Stop()

	var tickerCh <-chan time.Time
	if stat.duration > 0 {
		ticker := time.NewTicker(stat.duration)
		defer ticker.Stop()
		tickerCh = ticker.C
	}

	dirty := false
	syncShadow := func() {
		if dirty {
			shadow := make(map[string]int, len(stat.kv))
			for k, v := range stat.kv {
				shadow[k] = int(*v)
			}
			stat.shadowKv = shadow
		}
	}

outter:
	for {
		select {
		case item, ok := <-stat.ch:
			dirty = true
			if !ok {
				break outter
			}
			val, ok := stat.kv[item.key]
			if !ok {
				var n statVal
				val = &n
				stat.kv[item.key] = val
			}
			*val++
		case <-tickerCh:
			syncShadow()
			dirty = false
		}
	}

	syncShadow()
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
	buffer.WriteString("Stat[")

	if stat.shadowKv != nil {
		for k, v := range stat.shadowKv {
			buffer.WriteString(k)
			buffer.WriteString(":")
			buffer.WriteString(Itoa(v))
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
