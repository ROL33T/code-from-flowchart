package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	for i := 0; i <= 100; i++ {

		if i%2 == 0 {

			print("คู่", i)

		} else {
			print("คี่", i)
		}
	}
}
