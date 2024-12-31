// package main

// import "fmt"

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	input := "nlaebolko"
// 	fmt.Println(maxNumbersOfBalloons(input))
// }

// func maxNumbersOfBalloons(input string) int {
// 	b, a, l, o, n := 0, 0, 0, 0, 0
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == 'b' {
// 			b++
// 		} else if input[i] == 'a' {
// 			a++
// 		} else if input[i] == 'l' {
// 			l++
// 		} else if input[i] == 'o' {
// 			o++
// 		} else if input[i] == 'n' {
// 			n++
// 		}
// 	}

// 	l /= 2
// 	o /= 2

// 	return min(b, min(a, min(l, min(o, n))))
// }

// package main

// import "fmt"

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	input := "loobalnxballpoon"
// 	fmt.Println(maxNumbersOfBalloons(input))
// }

// func maxNumbersOfBalloons(input string) int {
// 	b, a, l, o, n := 0, 0, 0, 0, 0
// 	for i := 0; i < len(input); i++ {
// 		if input[i] == 'b' {
// 			b++
// 		} else if input[i] == 'a' {
// 			a++
// 		} else if input[i] == 'l' {
// 			l++
// 		} else if input[i] == 'o' {
// 			o++
// 		} else if input[i] == 'n' {
// 			n++
// 		}
// 	}
// 	l /= 2
// 	o /= 2

// 	return min(b, min(a, min(l, min(o, n))))
// }

package main

import "fmt"

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "dfrebjattljyyljgyogyoccn"
	fmt.Println(maxNumberOfBalloon(s))
}

func maxNumberOfBalloon(s string) int {
	b, a, l, o, n := 0, 0, 0, 0, 0
	for i := o; i < len(s); i++ {
		if s[i] == 'b' {
			b++
		} else if s[i] == 'a' {
			a++
		} else if s[i] == 'l' {
			l++
		} else if s[i] == 'o' {
			o++
		} else if s[i] == 'n' {
			n++
		}
	}
	l /= 2
	o /= 2

	return min(b, min(a, min(l, min(o, n))))
}
