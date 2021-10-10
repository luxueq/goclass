package week5

import "sync"

type Number struct {
	Buckets map[int64]float64
	Mutex   *sync.RWMutex
}

func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]float64),
		Mutex:   &sync.RWMutex{},
	}

	return r
}
