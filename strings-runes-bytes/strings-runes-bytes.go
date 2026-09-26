package main

import (
	"fmt"
	"strings"
)

func main() {
	// var myRune = 'a'
	var keyword = []string{"s", "u", "c", "k", "m", "y", "d", "i", "c", "k"}
	var strBuilder strings.Builder
	// var join string = ""

	for i := range keyword {
		strBuilder.WriteString(keyword[i])
	}
	// fmt.Println(join)
	var join = strBuilder.String()
	fmt.Printf("%v\n", join)
}
