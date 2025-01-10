// FINS PLAYERS WITH 0 AND 1 LOSES....

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	nums := [][]int{
// 		{1, 3},
// 		{2, 3},
// 		{3, 6},
// 		{5, 6},
// 		{5, 7},
// 		{4, 5},
// 		{4, 8},
// 		{4, 9},
// 		{10, 4},
// 		{10, 9},
// 	}
// 	fmt.Println(findWinners(nums))
// }

// func findWinners(nums [][]int) [][]int {
// 	lossessCount := make(map[int]int) // this map is stored loser count...

// 	for i := 0; i < len(nums); i++ {
// 		winner, loser := nums[i][0], nums[i][1] // there is only present 0 winners and 1 loser
// 		lossessCount[loser]++
// 		if _, exist := lossessCount[winner]; !exist { // if winner is not exist then also add winner with 0 losses
// 			lossessCount[winner] = 0
// 		}
// 	}

// 	// seperate pplayers based on loses...

// 	zeroLosses := []int{}
// 	oneLoss := []int{}

// 	keys := make([]int, 0, len(lossessCount))
// 	for k := range lossessCount {
// 		keys = append(keys, k)
// 	}

// 	for i := 0; i < len(keys); i++ {
// 		player := keys[i]
// 		losses := lossessCount[player]
// 		if losses == 0 {
// 			zeroLosses = append(zeroLosses, player)
// 		} else if losses == 1 {
// 			oneLoss = append(oneLoss, player)
// 		}
// 	}

// 	sort.Ints(zeroLosses)
// 	sort.Ints(oneLoss)

// 	return [][]int{zeroLosses, oneLoss}
// }

// Loser count in hashMap which name is lossessCount
// {
//     1: 0,
//     2: 0,
//     3: 2,
//     4: 1,
//     5: 1,
//     6: 2,
//     7: 1,
//     8: 1,
//     9: 2,
//     10: 0,
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	nums := [][]int{
// 		{1, 3},
// 		{2, 3},
// 		{3, 6},
// 		{5, 6},
// 		{5, 7},
// 		{4, 5},
// 		{4, 8},
// 		{4, 9},
// 		{10, 4},
// 		{10, 9},
// 	}
// 	fmt.Println(findWinners(nums))
// }

// func findWinners(nums [][]int) [][]int {
// 	lossessCount := make(map[int]int)

// 	for i := 0; i < len(nums); i++ {
// 		winner, loser := nums[i][0], nums[i][1]
// 		lossessCount[loser]++
// 		if _, exist := lossessCount[winner]; !exist {
// 			lossessCount[winner] = 0
// 		}
// 	}

// 	zeroLossess := []int{}
// 	oneLoss := []int{}

// 	keys := make([]int, 0, len(lossessCount))
// 	for k := range lossessCount {
// 		keys = append(keys, k)
// 	}

// 	for i := 0; i < len(keys); i++ {
// 		player := keys[i]
// 		losses := lossessCount[player]
// 		if losses == 0 {
// 			zeroLossess = append(zeroLossess, player)
// 		} else if losses == 1 {
// 			oneLoss = append(oneLoss, player)
// 		}
// 	}

// 	sort.Ints(zeroLossess)
// 	sort.Ints(oneLoss)

// 	return [][]int{zeroLossess, oneLoss}
// }

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{
		{1, 3},
		{2, 3},
		{3, 6},
		{5, 6},
		{5, 7},
		{4, 5},
		{4, 8},
		{4, 9},
		{10, 4},
		{10, 9},
	}
	fmt.Println((findWinners(nums)))
}

func findWinners(nums [][]int) [][]int {
	lossessCount := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		winner, loser := nums[i][0], nums[i][1]
		lossessCount[loser]++
		if _, exist := lossessCount[winner]; !exist {
			lossessCount[winner] = 0
		}
	}

	zeroLosses := []int{}
	oneLoss := []int{}

	keys := make([]int, 0, len(lossessCount))
	for k := range lossessCount {
		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {
		player := keys[i]
		losses := lossessCount[player]
		if losses == 0 {
			zeroLosses = append(zeroLosses, player)
		} else if losses == 1 {
			oneLoss = append(oneLoss, player)
		}
	}

	sort.Ints(zeroLosses)
	sort.Ints(oneLoss)

	return [][]int{zeroLosses, oneLoss}
}
