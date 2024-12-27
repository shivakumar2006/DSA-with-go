package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 4, 5}
	fmt.Println(removeDuplicate(nums))
}

func removeDuplicate(nums []int) int {
	unique := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			unique += 1
			nums[unique-1] = nums[i]
		}
	}
	return unique
}

// another way of solving this with the output len of nums as well as array

// package main

// import "fmt"

// func main() {
// 	nums := []int{1, 1, 2, 3, 4, 4, 5}
// 	length := removeDuplicate(nums)
// 	fmt.Println("total numbers : ", length)
// 	fmt.Println("unique elements : ", nums[:length])
// }

// func removeDuplicate(nums []int) int {
// 	unique := 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] != nums[i-1] {
// 			unique += 1
// 			nums[unique-1] = nums[i]
// 		}
// 	}
// 	return unique
// }
