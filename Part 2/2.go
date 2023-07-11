package main

import (
	"fmt"
)

func println(a ...interface{}) {
	fmt.Println(a...)
}

func print(a ...interface{}) {
	fmt.Print(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {

	var ARRAY_NUMBER_EVEN []int
	var ARRAY_NUMBER_ELEVEN []int
	var ARRAY_NUMBER_ZERO []int
	for i := 0; i <= 999; i++ {

		if i%2 == 0 {
			ARRAY_NUMBER_EVEN = append(ARRAY_NUMBER_EVEN, i)
			//print("คู่", i)
		}

		if i%7 == 0 {
			ARRAY_NUMBER_ELEVEN = append(ARRAY_NUMBER_ELEVEN, i)
		}

		if i%10 == 0 {
			ARRAY_NUMBER_ZERO = append(ARRAY_NUMBER_ZERO, i)
		}

	}

	println(ARRAY_NUMBER_EVEN)
	println(ARRAY_NUMBER_ELEVEN)
	println(ARRAY_NUMBER_ZERO)

}
