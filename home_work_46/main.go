package main

import (
	"fmt"
)

/**
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

var (
	res  [][]int //保存结果
	path []int   //保存路径
	st   []bool
)

func permute(nums []int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, len(nums))
	st = make([]bool, len(nums))

	backtrack(nums, 0)
	return res
}

func backtrack(output []int, cur int) {
	if cur == len(output) {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
	}
	for i := 0; i < len(output); i++ {
		if !st[i] {
			path = append(path, output[i])
			st[i] = true
			backtrack(output, cur+1)
			st[i] = false
			path = path[:len(path)-1]

		}
	}
}
