package main

import (
	"fmt"
	"strconv"
)

func noCharLen(s string) int {
	return len(s)
}

func main() {
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
	number1 := 17
	number2 := 24
	resultMessage := "The number is odd"

	// check if number is odd using modulus reminder
	if number1%2 == 0 {
		fmt.Println("The number is even")
	} else if number1%2 == 1 {
		fmt.Println(resultMessage)
	}

	fmt.Println("The # of chars are: ", noCharLen(strconv.Itoa(number2)))
}
