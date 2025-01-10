// // check if all number have equal number of occurrence

// package main

// import "fmt"

// func main() {
// 	s := "abcabc"
// 	fmt.Println(areEqualOccurrence(s))
// }

// func areEqualOccurrence(s string) bool {
// 	charSet := [26]int{}
// 	for i := 0; i < len(s); i++ {
// 		charSet[s[i]-'a']++
// 	}

// 	x := 0
// 	for i := 0; i < len(charSet); i++ {
// 		if charSet[i] != 0 {
// 			if x == 0 {
// 				x = charSet[i]
// 			} else if charSet[i] != x {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// package main

// import "fmt"

// func main() {
// 	s := "abacbc"
// 	fmt.Println(areEqualOccurrence(s))
// }

// func areEqualOccurrence(s string) bool {
// 	charSet := [26]int{}
// 	x := 0

// 	for i := 0; i < len(s); i++ {
// 		charSet[s[i]-'a']++
// 	}

// 	for i := 0; i < len(charSet); i++ {
// 		if charSet[i] != 0 {
// 			if x == 0 {
// 				x = charSet[i]
// 			} else if charSet[i] != x {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// package main

// import "fmt"

// func main() {
// 	s := "abacbc"
// 	fmt.Println(areEqualOccurrence(s))
// }

// func areEqualOccurrence(s string) bool {
// 	charSet := [26]int{}
// 	x := 0
// 	for i := 0; i < len(s); i++ {
// 		charSet[s[i]-'a']++
// 	}
// 	for i := 0; i < len(charSet); i++ {
// 		if charSet[i] != 0 {
// 			if x == 0 {
// 				x = charSet[i]
// 			} else if charSet[i] != x {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// package main

// import "fmt"

// func main() {
// 	s := "abcabc"
// 	fmt.Println(occurrence(s))
// }

// func occurrence(s string) bool {
// 	charSet := [26]int{}
// 	x := 0
// 	for i := 0; i < len(s); i++ {
// 		charSet[s[i]-'a']++
// 	}
// 	for i := 0; i < len(charSet); i++ {
// 		if charSet[i] != 0 {
// 			if x == 0 {
// 				x = charSet[i]
// 			} else if x != charSet[i] {
// 				return false
// 			}
// 		}
// 	return true
// }

package main

import "fmt"

func main() {
	s := "abcacb"
	fmt.Println(occurrence(s))
}

func occurrence(s string) bool {
	charSet := [26]int{}
	x := 0
	for i := 0; i < len(s); i++ {
		charSet[s[i]-'a']++
	}

	for i := 0; i < len(charSet); i++ {
		if charSet[i] != 0 {
			if x == 0 {
				x = charSet[i]
			} else if charSet[i] != x {
				return false
			}
		}
	}
	return true
}
