package main

import (
	"fmt"
	"sync"
	"time"
)

//编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

var channel chan int
var wg sync.WaitGroup

func main() {
	channel = make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			fmt.Println("写入数据: ", i)
			time.Sleep(1 * time.Second)
			channel <- i
		}
		close(channel)
	}()

	go func() {
		defer wg.Done()
		for {
			num, ok := <-channel
			if !ok {
				break
			}
			fmt.Println("读取数据: ", num)
		}
	}()

	wg.Wait()
}
