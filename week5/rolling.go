package week5

import (
	"sync"
	"time"
)

type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Tally float64
}

func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}

	return r
}

func (r *Number) getCurrentBucket() *numberBucket {
	time := time.Now().Unix()
	var bucket *numberBucket
	var ok bool

	if bucket, ok = r.Buckets[time]; !ok {
		bucket = &numberBucket{}
		r.Buckets[time] = bucket
	}

	return bucket
}

func(r *Number) removeOldBuckets() {
	minBucketTime := time.Now().Unix() - 10

	for timestamp := range r.Buckets {
		if timestamp <= minBucketTime {
			delete(r.Buckets, timestamp)
		}
	}
}


func (r *Number) IncreNumber(step float64) {
	if step == 0 {
		return
	}

	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	currentBucket := r.getCurrentBucket()

	currentBucket.Tally += step

	r.removeOldBuckets()
}