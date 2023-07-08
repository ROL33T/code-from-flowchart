package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	a := 0 //นับเลขคู่
	b := 0 //นับเลขคี่

	for i := 0; i <= 100; i++ {

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
}
