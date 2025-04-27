package main

import "fmt"

func main() {
	//nums := []int{2, 2, 1}
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
	var m = make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
