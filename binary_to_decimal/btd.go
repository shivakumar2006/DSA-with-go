package main

import (
	"fmt"
	"strconv"
)

func main() {
	binaryString := "111001000101111"
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

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	binaryString := "11010100110"
// 	decimalValue, err := strconv.ParseInt(binaryString, 2, 0)
// 	if err != nil {
// 		fmt.Println("Error : ", err)
// 	}
// 	fmt.Println(decimalValue)
// }
