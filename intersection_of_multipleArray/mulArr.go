// intersection of multiple array

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
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < len(nums[i]); j++ {
			hashMap[nums[i][j]]++
		}
	}
	result := []int{}
	for nums, count := range hashMap {
		if count == 0 {
			result = append(result, nums)
		}
	}
	return result
}
