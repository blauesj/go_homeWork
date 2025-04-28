package main

import "slices"

func main() {

}

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
	result := make([][]int, 0)
	e := intervals[0]
	result = append(result, e)
	for i := 1; i < len(intervals); i++ {
		if e[1] >= intervals[i][0] {
			e[0] = min(e[0], intervals[i][0])
			e[1] = max(e[1], intervals[i][1])
		} else {
			e = intervals[i]
			result = append(result, e)
		}
	}
	return result
}
