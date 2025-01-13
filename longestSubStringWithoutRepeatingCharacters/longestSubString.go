// given a string s. find the length of the longest substring without repeaing characters.

// package main

// import "fmt"

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	s := "abcabcbb"
// 	fmt.Println(longestSubString(s))
// }

// func longestSubString(s string) int {
// 	hashMap := make(map[byte]int)

// 	maxLength := 0
// 	start := 0

// 	for i := 0; i < len(s); i++ {
// 		if index, found := hashMap[s[i]]; found && index >= start {
// 			start = index + 1
// 		}

// 		hashMap[s[i]] = i

// 		maxLength = max(maxLength, i-start+1)
// 	}
// 	return maxLength
// }

// package main

// import "fmt"

// func max(a int, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	s := "abcabcbb"
// 	fmt.Println(longestSubString(s))
// }

// func longestSubString(s string) int {
// 	hashMap := make(map[byte]int)

// 	maxLength := 0
// 	start := 0

// 	for i := 0; i < len(s); i++ {
// 		if index, found := hashMap[s[i]]; found && index >= start {
// 			start = index + 1
// 		}

// 		hashMap[s[i]] = i

// 		maxLength = max(maxLength, i-start+1)
// 	}
// 	return maxLength
// }

package main

import "fmt"

func main() {
	s := "abcabcbb"
	fmt.Println(longestSubString(s))
}

func longestSubString(s string) int {
	hashMap := make(map[byte]int)

	maxlenght := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if index, found := hashMap[s[i]]; found && index >= start {
			start = index + 1
		}
	}
}
