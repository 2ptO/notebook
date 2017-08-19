package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
 * Go Routines examples with Mutex and Atomic Add
 */

func main() {
	fmt.Println("Go OS:", runtime.GOOS)
	fmt.Println("Go Num CPUs:", runtime.NumCPU())
	const numGoRoutines = 100
	var wg sync.WaitGroup
	var mu sync.Mutex

	//Incrementing counter with Mutex
	counter := int64(0) //why int64? because, atomic AddInt64 requires int64
	wg.Add(numGoRoutines)
	for i := 0; i < numGoRoutines; i++ {
		go func(id int) {
			mu.Lock()
			runtime.Gosched() //yields the CPU to other Go Routines to Run. Doesn't suspend
			counter++
			mu.Unlock()
			fmt.Printf("Go(%d) NumGoRoutines %d\n", id, runtime.NumGoroutine())
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Counter With Mutex: %d\n", counter)

	//Incrementing counter with Atomic Add
	wg.Add(numGoRoutines)
	for i := 0; i < numGoRoutines; i++ {
		go func(id int) {
			atomic.AddInt64(&counter, 1)
			fmt.Printf("Go(%d) Counter (%d) \n", id, atomic.LoadInt64(&counter))
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Printf("Counter with Atomic: %d\n", counter)
}
