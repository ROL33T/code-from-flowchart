package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	var count_number int
	print("กรุณาใส่จำนวนเลขที่ต้องการให้นับ: ")
	Scan(&count_number)

	a := 0 //นับเลขคู่
	b := 0 //นับเลขคี่
	if count_number > 0 {
		for i := 0; i <= count_number; i++ {

			if i != 0 {

				if i%2 == 0 {
					a++
					print("คู่", i)

				} else {
					b++
					print("คี่", i)
				}

			} else {
				print("เลข 0 จะไม่นับ")
			}

		}

		printf("เลขคู่ทั้งหมด %d เลขคี่ทั้งหมด %d", a, b)
	} else {
		print("กรุณาใส่จำนวนเลขที่ต้องการให้นับ 0 ขึ้นไป")
	}
}
