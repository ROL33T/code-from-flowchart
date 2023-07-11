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
	//หาต้องการจำแนกตัวเลข 0-999 เป็นarray 3ชุด โดยชุดแรกเก็บเลขคู่,ชุดที่สองเป็นเลขที่หาร7ลงตัว,ชุดที่สามเป็นตัวเลขที่ลงท้ายด้วย0เท่านั้น หลังจำแนกข้อมูลให้แสดงข้อมูลทั้ง3ออกมา

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

	println("เลขคู่ทั้งหมด : \n", ARRAY_NUMBER_EVEN)
	println("เลข7ที่หารลงตัวทั้งหมด : \n", ARRAY_NUMBER_ELEVEN)
	println("หาเลขที่ลงท้ายด้วย 0ตัวทั้งหมด : \n", ARRAY_NUMBER_ZERO)

}
