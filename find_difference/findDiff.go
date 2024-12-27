// package main

// import "fmt"

// func main() {
// 	n := 10
// 	m := 2
// 	fmt.Println((findDifference(n, m)))
// }

// func findDifference(n int, m int) int {
// 	x := n / m
// 	num2 := m * x * (x + 1) / 2
// 	num1 := n * (n + 1) / 2
// 	return num1 - num2
// }

// another way of solve this through loop

package main

import "fmt"

func main() {
	n := 10
	m := 2
	fmt.Println(findDifference(n, m))
}

func findDifference(n int, m int) int {
	num1 := 0
	num2 := 0
	for i := 0; i <= n; i++ {
		num1 += i
	}
	for i := m; i <= n; i += m {
		num2 += i
	}
	return num1 - num2
}
