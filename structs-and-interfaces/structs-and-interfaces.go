package main

import "fmt"

type human struct {
	name               string
	height             float32
	age                int
	hands              int
	legs               int
	eyes               int
	mouth              int
	relationshipStatus string
}

func main() {
	var seunName human = human{name: "seun", height: 1.76, age: 23, hands: 2, legs: 2, eyes: 2, mouth: 1, relationshipStatus: "single"}

	fmt.Printf("my name is %v i am %v years old\n", seunName.name, seunName.age)
}
