package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int{18, 69, 70, 58, 42, 35, 110, 99}
	fmt.Println(binarySearch(array, 69))
}

func binarySearch(array []int, x int) int {
	l := 0
	r := len(array) - 1
	sort.Ints(array)
	for l <= r {
		mid := (l + r) / 2
		if array[mid] == x {
			return mid
		} else if x < array[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
