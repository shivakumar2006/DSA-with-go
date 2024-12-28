package main

import (
	"fmt"
	"strconv"
)

func main() {
	binaryString := "0101"
	decimalValue, err := strconv.ParseInt(binaryString, 2, 0)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(decimalValue)
}

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	binaryString := "10010010010011"
// 	decimalValue, err := strconv.ParseInt(binaryString, 2, 0)
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 	}
// 	fmt.Println(decimalValue)
// }
