package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var sum int32 = 0
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(10)
	for range 10 {
		go func() {
			for range 1000 {
				atomic.AddInt32(&sum, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("sum: ", sum)

}
