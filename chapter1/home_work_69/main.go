package main

import "fmt"

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(10))

}

func mySqrt(x int) int {
	for i := 0; i <= x; i++ {
		if i*i == x {
			return i
		}
		if i*i > x {
			return i - 1
		}
	}
	return 0
}
