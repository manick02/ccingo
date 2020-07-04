package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		// this go routine will not be garbage collected here
		wg.Done()
		<-c
	}
	const numRoutines = 1e6
	wg.Add(numRoutines)
	before := memconsumed()
	for i := 0; i < numRoutines; i++ {
		go noop()
	}
	wg.Wait()
	after := memconsumed()
	fmt.Println()
	fmt.Printf("memory used per goroutine: %.3f kb\n", float64(after-before)/(numRoutines*1024))
}

func memconsumed() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}
