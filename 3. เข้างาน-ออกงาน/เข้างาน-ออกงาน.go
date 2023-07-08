package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	t1 := 9.30  //เวลาเข้างาน
	t2 := 15.59 //เวลาออกงาน

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
