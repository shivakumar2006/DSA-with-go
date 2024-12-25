package main

import "fmt"

func main() {
	array := []int{18, 32, 44, 58, 69, 72, 99, 110}
	fmt.Println(binarySearch(array, 69))
}

func binarySearch(array []int, x int) int {
	l := 0
	r := len(array) - 1
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
