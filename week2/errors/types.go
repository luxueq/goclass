package errors

import "sync"

func EmptyRow(reason string, message string) *Error {
	return Newf(110001, reason, message)
}

type testMutex struct {
	 mutex sync.Mutex
}

