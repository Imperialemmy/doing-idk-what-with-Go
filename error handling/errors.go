package main

import (
	"errors"
	"fmt"
)

func division(top int, bottom int) (int, int, error) {
	var err error
	if bottom == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result = top / bottom
	var remainder = top % bottom
	return result, remainder, err
}

func main() {
	var top int = 10
	var bottom int = 2
	var result, remainder, err = division(top, bottom)
	if err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Printf("The result of the division is %v remainder %v\n", result, remainder)
}
