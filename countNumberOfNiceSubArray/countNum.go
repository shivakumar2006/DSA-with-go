// COUNT NUMBER OF NICE SUBARRAY
// Given an array num and an integer k. A continuous subarray is called NICE if there are k odd numbers on it.

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 2, 1, 1}
// 	k := 3
// 	fmt.Println(numberOfSubArray(nums, k))
// }

// func numberOfSubArray(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]%2 != 0 {
// 			current++
// 		}
// 		if count, exist := hashMap[current-k]; exist {
// 			result += count
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 2, 1, 1}
// 	k := 3
// 	fmt.Println(numberOfSubArray(nums, k))
// }

// func numberOfSubArray(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]%2 != 0 {
// 			current++
// 		}
// 		if count, exist := hashMap[current-k]; exist {
// 			result += count
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 2, 1, 1}
// 	k := 3
// 	fmt.Println(numberOfSubArray(nums, k))
// }

// func numberOfSubArray(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]%2 != 0 {
// 			current++
// 		}
// 		if value, exist := hashMap[current-k]; exist {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 2, 1, 1}
// 	k := 3
// 	fmt.Println(numberOfSubArray(nums, k))
// }

// func numberOfSubArray(nums []int, k int) int {
// 	result := 0
// 	curr := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i]%2 != 0 {
// 			curr++
// 		}
// 		if value, exist := hashMap[curr-k]; exist {
// 			result += value
// 		}
// 		hashMap[curr]++
// 	}
// 	return result
// }

package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 1, 1}
	k := 3
	fmt.Println(numberOfSubArray(nums, k))
}

func numberOfSubArray(nums []int, k int) int {
	result := 0
	current := 0
	hashMap := make(map[int]int)
	hashMap[0] = 1
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 != 0 {
			current++
		}
		if value, exist := hashMap[current-k]; exist {
			result += value
		}
		hashMap[current]++
	}
	return result
}
