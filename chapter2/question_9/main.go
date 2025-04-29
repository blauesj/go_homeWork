package main

import (
	"fmt"
	"sync"
)

var sum int = 0
var wg sync.WaitGroup
var lock sync.Mutex

func main() {
	wg.Add(10)
	for range 10 {
		go func() {
			for range 1000 {
				lock.Lock()
				sum += 1
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("sum: ", sum)

}
