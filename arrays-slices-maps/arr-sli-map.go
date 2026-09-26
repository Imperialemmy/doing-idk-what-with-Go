package main

import (
	"fmt"
)

// func main() {
// 	ages := make(map[string]int)
// 	ages["John"] = 25
// 	ages["Mary"] = 30
// 	for key, value := range ages {
// 		fmt.Println(key, value)
// 	}
// }

func displayages(name string, age int) {
	ages := map[string]int{
		name: age,
	}
	for key, value := range ages {
		fmt.Println(key, value)
	}
}

func main() {
	// displayages("seun", 23)
	// displayages("seun girl", 18)
	// displayages("daddy", 62)
	// displayages("mummy", 57)
	// displayages("ay", 26)
	displayages("johnny", 23)
	displayages("kiing", 24)
	fmt.Println("omo una don old o")
}
