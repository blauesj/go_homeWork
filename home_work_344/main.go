package main

import "fmt"

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))

}

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

}
