package test

import (
	"testing"

	"github.com/sweetycode/goal"
)

func BenchmarkMapRW(b *testing.B) {
	kv := map[string]int{
		"key": 1,
	}
	for i := 0; i < b.N; i++ {
		kv["key"]++
	}
}

func BenchmarkPointerMap(b *testing.B) {
	kv := map[string]*goal.StatVal{
		"key": &goal.StatVal{},
	}

	for i := 0; i < b.N; i++ {
		kv["key"].Incr(1)
	}
}
