// max_sum_of_pair_with_equal_sum_of_digit

package main

import "fmt"

func main() {
	array := []int{18, 43, 36, 13, 7}
	fmt.Println(maximumCard(array))
}

func maximumCard(array []int) int {
	hashMap := make(map[int]int)
	result := -1
	for i := 0; i < len(array); i++ {
		s := sum(array[i])
		if value, ok := hashMap[s]; ok {
			result = max(result, array[i]+value)
			if array[i] > value {
				hashMap[s] = array[i]
			}
		} else {
			hashMap[s] = array[i]
		}
	}
	return result
}

func sum(n int) int {
	target := 0
	for n > 0 {
		target += n % 10
		n = n / 10
	}
	return target
}