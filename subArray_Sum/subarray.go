// subarray sum equal to K

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 1}
// 	k := 2
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		current += nums[i]
// 		if value, found := hashMap[current-k]; found {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// Input: nums = [1, 1, 1], k = 2

// Initialization:

// result = 0
// current = 0
// hashMap = {0: 1}
// Iteration 1 (i = 0, nums[0] = 1):

// current = current + nums[0] = 0 + 1 = 1
// current - k = 1 - 2 = -1 (not in hashMap)
// Update hashMap: {0: 1, 1: 1}
// Iteration 2 (i = 1, nums[1] = 1):

// current = current + nums[1] = 1 + 1 = 2
// current - k = 2 - 2 = 0 (exists in hashMap with value 1)
// Increment result to 1 (result = 1)
// Update hashMap: {0: 1, 1: 1, 2: 1}
// Iteration 3 (i = 2, nums[2] = 1):

// current = current + nums[2] = 2 + 1 = 3
// current - k = 3 - 2 = 1 (exists in hashMap with value 1)
// Increment result to 2 (result = 2)
// Update hashMap: {0: 1, 1: 1, 2: 1, 3: 1}
// Final Output: 2.

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 1}
// 	k := 2
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	curr := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		curr += nums[i]
// 		if val, found := hashMap[curr-k]; found {
// 			result += val
// 		}
// 		hashMap[curr]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 1}
// 	k := 2
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		current += nums[i]
// 		if value, found := hashMap[current-k]; found {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 3}
// 	k := 3
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		current += nums[i]
// 		if value, found := hashMap[current-k]; found {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 3}
// 	k := 3
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		current += nums[i]
// 		if value, found := hashMap[current-k]; found {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int {1, 1, 1}
// 	k := 2;
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	curr := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		curr += nums[i]
// 		if value, found := hashMap[curr - k]; found {
// 			result += value
// 		}
// 		hashMap[curr]++
// 	}
// 	return result
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 2, 3}
// 	k := 3
// 	fmt.Println(subArraySum(nums, k))
// }

// func subArraySum(nums []int, k int) int {
// 	result := 0
// 	current := 0
// 	hashMap := make(map[int]int)
// 	hashMap[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		current += nums[i]
// 		if value, found := hashMap[current-k]; found {
// 			result += value
// 		}
// 		hashMap[current]++
// 	}
// 	return result
// }

package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	k := 3
	fmt.Println(subArraySum(nums, k))
}

func subArraySum(nums []int, k int) int {
	result := 0
	current := 0
	hashMap := make(map[int]int)
	hashMap[0] = 1
	for i := 0; i < len(nums); i++ {
		current += nums[i]
		if value, found := hashMap[current-k]; found {
			result += value
		}
		hashMap[current]++
	}
	return result
}
