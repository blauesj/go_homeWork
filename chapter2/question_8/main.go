package main

import (
	"fmt"
	"time"
)

var channel chan int

func main() {
	channel = make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			fmt.Println("写入数据: ", i)
			time.Sleep(1 * time.Second)
			channel <- i
		}
	}()
	for {
		num, ok := <-channel
		if !ok {
			break
		}
		fmt.Println("读取数据: ", num)
	}
}
