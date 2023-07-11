package main

import (
	"fmt"
	"sort"
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
	//จาก array ที่กำหนดจงเรียงค่าในarrayจากน้อยไปมากและจากมากไปน้อยแล้วให้แสดง array นั้นออกมา
	x := []int{1, 5, 6, 9, 10, 11, 25, 99, 0, -5}
	sort.Ints(x)

	println(x)
}
