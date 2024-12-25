// package main

// import "fmt"

// func main() {
// 	array := []int{1, 2, 3, 4, 5}
// 	fmt.Println(rotateArray(array, 3))
// }

// func rotateArray(array []int, d int) []int {
// 	n := len(array)
// 	rev(array, 0, d-1)
// 	rev(array, d, n-1)
// 	rev(array, 0, n-1)
// 	return array
// }

// func rev(array []int, l int, r int) {
// 	for l < r {
// 		temp := array[l]
// 		array[l] = array[r]
// 		array[r] = temp
// 		l++
// 		r--
// 	}
// }

package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Println(rotateArray(array, 3))
}

func rotateArray(array []int, d int) []int {
	n := len(array)
	rev(array, 0, d-1)
	rev(array, d, n-1)
	rev(array, 0, n-1)
	return array
}

func rev(array []int, l int, r int) {
	for l < r {
		temp := array[l]
		array[l] = array[r]
		array[r] = temp
		l++
		r--
	}
}
