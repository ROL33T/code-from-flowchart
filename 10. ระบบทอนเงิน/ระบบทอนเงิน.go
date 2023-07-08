package main

import (
	"fmt"
)

func print(a ...interface{}) {
	fmt.Println(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	m := 0 //จำนวนเงินในกระเป๋า
	total := 0 // จำนวนธนบัตรและเหรียญทั้งหมด
	if m > 0 {
		if m >= 1000 {
			total += m / 1000
			m %= 1000
		}

		if m >= 500 {
			total++
			m -= 500
		}

		if m >= 100 {
			total += m / 100
			m %= 100
		}

		if m >= 10 {
			total += m / 10
			m %= 10
		}

		if m >= 5 {
			total += m / 5
			m %= 5
		}

		total += m // นับจำนวนเหรียญ 1
		print("จำนวนธนบัตรและเหรียญทั้งหมด:", total) 
	} else {
		total = 0
		print("ไม่สามารถคำนวนได้เนื่องจากคุณไม่มีเงินให้ทอน")
	}
	
}
