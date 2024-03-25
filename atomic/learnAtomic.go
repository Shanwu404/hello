package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func AtomicAdd() {
	// var a int32 =  0
	var a atomic.Value
	a.Store(0)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// atomic.AddInt32(&a, 1)
			fmt.Println(a)
		}()
	}
	wg.Wait()
	timeSpends := time.Since(start).Nanoseconds()
	fmt.Printf("use atomic a is %d, spend time: %v\n", a.Load(), timeSpends)
}

func main() {
	AtomicAdd()
}
