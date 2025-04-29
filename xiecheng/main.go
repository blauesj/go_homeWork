package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	go test()
	for i := 0; i < 10; i++ {
		fmt.Println("hello world ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello golang ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
