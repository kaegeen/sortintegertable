package main

import (
	"fmt"
)

func main() {
	s := []int{5, 4, 3, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
