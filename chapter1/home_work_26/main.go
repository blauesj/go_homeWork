package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}      // 输入数组
	expectedNums := []int{1, 2} // 长度正确的期望答案
	k := removeDuplicates(nums)
	if k != len(expectedNums) {
		panic("返回的长度与期望长度不相等")
	}
	for i := 0; i < len(expectedNums); i++ {
		if nums[i] != expectedNums[i] {
			panic("返回的数组与期望数组不相等")
		}
	}
	fmt.Println("通过测试")

}

func removeDuplicates(nums []int) int {
	i1, i2 := 0, 1
	for i2 < len(nums) {
		if nums[i1] != nums[i2] {
			i1++
			nums[i1] = nums[i2]
		}
		i2++
	}
	return i1 + 1

}
