package main

import (
	"fmt"
)

func main() {
	array := []int{5, 6, 7, 8, 9, 10, 1, 2, 3}
	key := 9
	fmt.Println(search(array, 0, len(array)-1, key))
}

func search(array []int, l int, r int, key int) int {
	pivot := getPivot(array, l, r)
	if pivot == -1 {
		return binarySearch(array, l, r, key)
	} else if array[pivot] == key {
		return pivot
	} else if array[l] <= key {
		return binarySearch(array, l, pivot-1, key)
	}
	return binarySearch(array, pivot+1, r, key)
}

func getPivot(array []int, l int, r int) int {
	for l <= r {
		mid := (l + r) / 2
		if mid < r && array[mid] > array[mid+1] {
			return mid
		} else if mid > l && array[mid-1] > array[mid] {
			return mid - 1
		} else if array[l] <= array[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func binarySearch(array []int, l int, r int, x int) int {
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

// package main

// import "fmt"

// func main() {
// 	array := []int{4, 5, 6, 7, 8, 9, 10, 1, 2, 3}
// 	key := 9
// 	fmt.Println(search(array, 0, len(array)-1, key))
// }

// func search(array []int, l int, r int, key int) int {
// 	pivot := getPivot(array, l, r)
// 	if pivot == -1 {
// 		return binarySearch(array, l, r, key)
// 	} else if array[pivot] == key {
// 		return pivot
// 	} else if array[l] <= key {
// 		return binarySearch(array, l, pivot-1, key)
// 	}
// 	return binarySearch(array, pivot+1, r, key)
// }

// func getPivot(array []int, l int, r int) int {
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if mid < r && array[mid] > array[mid+1] {
// 			return mid
// 		} else if mid > l && array[mid-1] > array[mid] {
// 			return mid - 1
// 		} else if array[l] <= array[mid] {
// 			l = mid + 1
// 		} else {
// 			r = mid - 1
// 		}
// 	}
// 	return -1
// }

// func binarySearch(array []int, l int, r int, x int) int {
// 	for l <= r {
// 		mid := (l + r) / 2
// 		if array[mid] == x {
// 			return mid
// 		} else if x < array[mid] {
// 			r = mid - 1
// 		} else {
// 			l = mid + 1
// 		}
// 	}
// 	return -1
// }
