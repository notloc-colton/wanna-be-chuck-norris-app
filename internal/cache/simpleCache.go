//TODO: This package is a stubbed out replacement for a distributed cache (such as Redis, memcached, etc.)

package cache

import (
	"sync"
)

type SimpleCache[T any] interface {
	Pop() (T, bool)
	Add(newItem T)
}
type simpleCache[T any] struct {
	data []T
	mu   sync.Mutex
}

func (sc *simpleCache[T]) Pop() (T, bool) {
	if len(sc.data) <= 0 {
		return *new(T), false
	}
	sc.mu.Lock()
	defer sc.mu.Unlock()
	top := sc.data[0]
	sc.data = sc.data[1:]
	return top, true
}
func (sc *simpleCache[T]) Add(newItem T) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.data = append(sc.data, newItem)
}

func NewSimpleCache[T any]() *simpleCache[T] {
	return &simpleCache[T]{
		data: []T{},
	}
}
