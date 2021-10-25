package main

import (
	"sync"
	"testing"
)

var bytePool = sync.Pool{
	New: newPool,
}

func newPool() interface{} {
	b := make([]byte, 1024)
	return &b
}
func BenchmarkAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := make([]byte, 1024)
		_ = obj
	}
}

func BenchmarkPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := bytePool.Get().(*[]byte)
		_ = obj
		bytePool.Put(obj)
	}
}
