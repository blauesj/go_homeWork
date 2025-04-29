package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*
*编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
 */
var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("routine 1,print number: ", strconv.Itoa(i))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("routine 2,print number: ", strconv.Itoa(i))
		}
	}()

	wg.Wait()
}
