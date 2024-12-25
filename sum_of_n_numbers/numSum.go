package main

import "fmt"

func main() {
	n := 10
	fmt.Println(sum(n))
}

func sum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum
}
