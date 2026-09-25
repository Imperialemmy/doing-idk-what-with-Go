package main

import (
	"fmt"
)

// func division(top int, bottom int) (int, int, error) {
// 	var err error
// 	if bottom == 0 {
// 		err = errors.New("Cannot Divide by Zero ")
// 		return 0, 0, err
// 	}
// 	var result = top / bottom
// 	var remainder = top % bottom
// 	return result, remainder, err
// }

// func main() {
// 	var top int = 10
// 	var bottom int = 2
// 	var result, remainder, err = division(top, bottom)
// 	//if, else and elseif method
// 	// if err != nil {
// 	// 	fmt.Printf(err.Error())
// 	// } else {
// 	// 	fmt.Printf("The result of the division is %v remainder %v\n", result, remainder)
// 	// }
// 	//switch case method
// 	switch {
// 	case err != nil:
// 		fmt.Printf(err.Error())
// 	case remainder == 0:
// 		fmt.Printf("the result of the division is %v\n", result)
// 	default:
// 		fmt.Printf("the result of the the division is %v remainder %v\n", result, remainder)
// 	}
// }

func multiplication(number1 int, number2 int) int {
	var result = number1 * number2
	return result
}

func main() {
	var number1 int = 10000
	var number2 int = 10000
	var result = multiplication(number1, number2)

	fmt.Printf("the result of the multiplication is %v\n", result)

}

// switches cases automatically apply breaks
