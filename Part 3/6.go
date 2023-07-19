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

	println("กรุณาใส่ Input One |  : ")
	input_one := 0
	Scan(&input_one)

	println("กรุณาใส่ Input Two |  : ")
	input_two := 0
	Scan(&input_two)

	if input_one > input_two {
		println("ใส่ตัวเลขไม่ถูกต้อง")
		return
	}

	for i := input_one; i <= input_two; i++ {
		if i%7 == 0 {
			printf("%d, ", i)
		}
	}

	println()

}
