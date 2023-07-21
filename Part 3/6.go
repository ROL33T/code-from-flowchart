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

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	var input_one, input_two int64

	println("กรุณาใส่ Input One |  : ")
	if _, err := fmt.Scan(&input_one); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		return
	}

	println("กรุณาใส่ Input Two |  : ")
	if _, err := fmt.Scan(&input_two); err != nil {
		println("ไม่สามารถอ่านค่า Input Two ได้")
		return
	}

	if input_two > input_one {
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
