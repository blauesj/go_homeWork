package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
 */
var wg sync.WaitGroup

func main() {
	taskList = append(taskList, func() {
		fmt.Println("task1")
		time.Sleep(time.Second * time.Duration(randomInt(1, 5)))
	}, func() {
		fmt.Println("task2")
		time.Sleep(time.Second * time.Duration(randomInt(1, 5)))
	}, func() {
		fmt.Println("task3")
		time.Sleep(time.Second * time.Duration(randomInt(1, 5)))
	})

	for _, task := range taskList {
		wg.Add(1)
		go func(task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			end := time.Now()
			fmt.Println("task execute time: ", end.Sub(start))
		}(task)
	}
	wg.Wait()

}

var taskList = []func(){}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
