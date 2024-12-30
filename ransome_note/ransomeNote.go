// package main

// import "fmt"

// func main() {
// 	ransomeNote := "aa"
// 	magazine := "aab"
// 	fmt.Println(canConstruct(ransomeNote, magazine))
// }

// func canConstruct(ransomeNote string, magazine string) bool {
// 	charSet := [26]int{}

// 	for i := 0; i < len(magazine); i++ {
// 		charSet[magazine[i]-'a']++
// 	}

// 	for i := 0; i < len(ransomeNote); i++ {
// 		charSet[ransomeNote[i]-'a']--
// 		if charSet[ransomeNote[i]-'a'] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// package main

// import "fmt"

// func main() {
// 	ransomeNote := "aa"
// 	magazine := "aab"
// 	fmt.Println(canConstruct(ransomeNote, magazine))
// }

// func canConstruct(ransomeNote string, magazine string) bool {
// 	charSet := [26]int{}
// 	for i := 0; i < len(magazine); i++ {
// 		charSet[magazine[i]-'a']++
// 	}

// 	for i := 0; i < len(ransomeNote); i++ {
// 		charSet[ransomeNote[i]-'a']--
// 		if charSet[ransomeNote[i]-'a'] < 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

package main

import "fmt"

func main() {
	ransomeNote := "aa"
	magazine := "aab"
	fmt.Println(canConstruct(ransomeNote, magazine))
}

func canConstruct(ransomeNote string, magazine string) bool {
	charSet := [26]int{}

	for i := 0; i < len(magazine); i++ {
		charSet[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomeNote); i++ {
		charSet[ransomeNote[i]-'a']--
		if charSet[ransomeNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
