// FIND 2 NUMBERS IN THE GIVEN ARRAY WHOSE SUM IS = TARGET.
// USE TWO POINTER APROACH.

// package main

// import "fmt"

// func main() {
// 	nums := []int{2, 3, 5, 7, 11}
// 	target := 10
// 	result := twoSum(nums, target)
// 	if result == nil {
// 		fmt.Println("there is not any pair of 2 sum whose equals to the target")
// 	} else {
// 		fmt.Println("indices of number : ", result)
// 	}
// }

// func twoSum(nums []int, target int) []int {
// 	l := 0
// 	r := len(nums) - 1
// 	for l < r {
// 		if nums[l]+nums[r] < target {
// 			l++
// 		} else if nums[l]+nums[r] > target {
// 			r--
// 		} else {
// 			return []int{l, r}
// 		}
// 	}
// 	return nil
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{2, 3, 5, 7, 11}
// 	target := 10
// 	result := twoSum(nums, target)
// 	if result == nil {
// 		fmt.Println("there is not any two pair of array whose sum is equal to the target...")
// 	} else {
// 		fmt.Println("Indices of numbers : ", result)
// 	}
// }

// func twoSum(nums []int, target int) []int {
// 	l := 0
// 	r := len(nums) - 1
// 	for i := 0; i < len(nums); i++ {
// 		if nums[l]+nums[r] < target {
// 			l++
// 		} else if nums[l]+nums[r] > target {
// 			r--
// 		} else {
// 			return []int{l, r}
// 		}
// 	}
// 	return nil
// }

// package main

// import "fmt"

// func main() {
// 	nums := []int{2, 3, 5, 7, 11}
// 	target := 10
// 	result := twoSum(nums, target)
// 	if result == nil {
// 		fmt.Println("There is not any 2 numbers i spresent whose sum is equal to target...")
// 	} else {
// 		fmt.Println("Indices of numbers : ", result)
// 	}
// }

// func twoSum(nums []int, target int) []int {
// 	l := 0
// 	r := len(nums) - 1
// 	for l < r {
// 		if nums[l]+nums[r] < target {
// 			l++
// 		} else if nums[l]+nums[r] > target {
// 			r--
// 		} else {
// 			return []int{l, r}
// 		}
// 	}
// 	return nil
// }

package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 7, 11}
	target := 10
	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("there isnot any 2 numbers is present whose sum is equals to target...")
	} else {
		fmt.Println("Indices of numbers : ", result)
	}
}

func twoSum(nums []int, target int) []int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] < target {
			l++
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			return []int{l, r}
		}
	}
	return nil
}
