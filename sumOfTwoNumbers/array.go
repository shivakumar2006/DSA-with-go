// FIND 2 NUMBERS IN THE GIVEN ARRAY WHOSE SUM IS = TARGET.
// USE TWO POINTER APROACH.

package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 7, 11}
	target := 10
	result := twoSum(nums, target)
	if result == nil {
		fmt.Println("there is not any pair of 2 sum whose equals to the target")
	} else {
		fmt.Println("indices of number : ", target)
	}
}

func twoSum(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		if nums[l]+nums[r] < target {
			l++
		} else if nums[l]+nums[r] > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
