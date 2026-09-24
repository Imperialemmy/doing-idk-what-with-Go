package main

import "fmt"

// import "fmt"

// func main() {
// 	const myName string = "victor"
// 	fmt.Println("my name is " + myName)

// 	var myAge int16 = 23
// 	fmt.Println(myAge)
// }

func division(top int, bottom int) (int, int) {
	var result = top / bottom
	var remainder = top % bottom
	return result, remainder
}

func main() {
	var top int = 10
	var bottom int = 2
	var result, remainder int = division(top, bottom)

	fmt.Printf("The result of the division is %v remainder %v\n", result, remainder)
}
