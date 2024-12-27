// package main

// import "fmt"

// func main() {
// 	jewels := "aA"
// 	stones := "aAAbbbbb"
// 	fmt.Println(check(jewels, stones))
// }

// func check(jewels string, stones string) int {
// 	hashMap := make(map[byte]bool)

// 	for i := 0; i < len(jewels); i++ {
// 		hashMap[jewels[i]] = true
// 	}

// 	result := 0
// 	for i := 0; i < len(stones); i++ {
// 		if hashMap[stones[i]] {
// 			result++
// 		}
// 	}
// 	return result
// }

package main

import "fmt"

func main() {
	jewels := "aA"
	stones := "aAAbbbbb"
	fmt.Println(check(jewels, stones))
}

func check(jewels string, stones string) int {
	hashMap := make(map[byte]bool)
	for i := 0; i < len(jewels); i++ {
		hashMap[jewels[i]] = true
	}

	result := 0
	for i := 0; i < len(stones); i++ {
		if hashMap[stones[i]] {
			result++
		}
	}
	return result
}
