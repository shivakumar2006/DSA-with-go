// intersection of multiple array

// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	n := len(nums)
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}
// 	result := []int{}
// 	for nums, count := range hashMap {
// 		if count == n {
// 			result = append(result, nums)
// 		}
// 	}
// 	return result
// }

//DEBUG CODE
// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	n := len(nums)

// 	fmt.Println("input arrays : ", nums)

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	fmt.Println("hashMap : ", hashMap)

// 	result := []int{}
// 	for num, count := range hashMap {
// 		fmt.Printf("Number: %d, Count: %d \n", num, count)
// 		if count == n {
// 			result = append(result, num)
// 		}
// 	}

// 	fmt.Println("Result", result)

// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	result := []int{}
// 	for num, count := range hashMap {
// 		if count == len(nums) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	result := []int{}
// 	for num, count := range hashMap {
// 		if count == len(nums) {
// 			result = append(result, num)
// 		}
// 	}
// 	sort.Ints(result)
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	result := []int{}
// 	for num, count := range hashMap {
// 		if count == len(nums) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	result := []int{}
// 	for num, count := range hashMap {
// 		if count == len(nums) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := [][]int{
// 		{3, 1, 2, 4, 5},
// 		{1, 2, 3, 4},
// 		{3, 4, 5, 6},
// 	}
// 	fmt.Println(intersection(nums))
// }

// func intersection(nums [][]int) []int {
// 	hashMap := make(map[int]int)
// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums[i]); j++ {
// 			hashMap[nums[i][j]]++
// 		}
// 	}

// 	result := []int{}
// 	for num, count := range hashMap {
// 		if count == len(nums) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

package main

import "fmt"

func main() {
	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}
	fmt.Println(intersection(nums))
}

func intersection(nums [][]int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			hashMap[nums[i][j]]++
		}
	}

	result := []int{}
	for value, count := range hashMap {
		if count == len(nums) {
			result = append(result, value)
		}
	}
	return result
}
