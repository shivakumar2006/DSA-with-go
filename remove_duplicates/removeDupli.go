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
