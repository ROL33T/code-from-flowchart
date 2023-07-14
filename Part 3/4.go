package main

import "fmt"

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
	height := 15
	for i := 1; i <= height; i++ {
		if i == 1 || i == height/2+1 || i == height {

			println("******")

		} else {
			println("*    *")
		}
	}

}
