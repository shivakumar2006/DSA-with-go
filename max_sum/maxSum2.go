package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

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
	total := 0
	for n > 0 {
		total += n % 10
		n = n / 10
	}
	return total
}
