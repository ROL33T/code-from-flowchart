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

func Scanln(a ...interface{}) {
	fmt.Scanln(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	var n int

	print("กรุณากรอกเลขที่ต้องการให้แสดงใน map: ")
	Scan(&n)

	data := make(map[int]int)

	for i := 1; i <= n; i++ {
		data[i] = i * i
	}

	println("Output:", data, "size of array:", len(data))

}
