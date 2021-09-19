package main

import "fmt"

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	a
	b
	mutexWaiterShift = iota
	starvationThresholdNs = 1e6
)

const (
	c = iota
)

func main()  {
	fmt.Printf("%d\n", mutexLocked)
	fmt.Printf("%d\n", mutexWoken)
	fmt.Printf("%d\n", mutexStarving)
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", mutexWaiterShift)
	fmt.Printf("%d\n", c)
	fmt.Printf("%.4f\n", starvationThresholdNs)
}
