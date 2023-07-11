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
	/*เขียนโค๊ดแสดงรูป
	        *
	      * * *
	    * * * * *
	  * * * * * * *
	* * * * * * * * *
	โดยใช้ Loop-for เท่านั้น*/

	height := 5

	for i := 1; i <= height; i++ {
		for j := 1; j <= height-i; j++ {
			print("  ")
		}

		for k := 1; k <= 2*i-1; k++ {
			print("* ")
		}

		println()
	}

}
