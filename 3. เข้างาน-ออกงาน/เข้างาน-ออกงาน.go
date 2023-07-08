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
	var t1 float64
	print("กรุณาใส่เวลาเข้างาน: ") //เวลาเข้างาน
	Scan(&t1)

	JOIN_WORK := false

	if t1 <= 9.00 {
		print("เข้างาน")
		JOIN_WORK = true

	} else if t1 >= 9.30 {
		print("ขาดงาน")
		JOIN_WORK = false
	} else {
		print("สาย")
		JOIN_WORK = true
	}

	if JOIN_WORK {

		var t2 float64
		print("กรุณาใส่เวลาออกงาน: ") //เวลาออกงาน
		Scan(&t2)

		if t2 <= 15.59 {
			print("ขาดงาน")
		} else {
			JOIN_WORK = true
		}

		if t2 >= 16.00 && t2 <= 18.00 {
			print("เลิกงาน")
		} else {
			if t2 > 18.00 {
				print("OT")
			}
		}

	}
}
