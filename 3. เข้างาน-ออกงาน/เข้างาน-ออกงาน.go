package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	t1 := 9.00 //เวลาเข้างาน
	t2 := 18.59 //เวลาออกงาน

	if t1 <= 9.00 {
		print("เข้างาน")
		if t2 <= 15.59 {
			print("ขาดงาน")
		} else {
			if t2 >= 16.00 && t2 <= 18.00 {
				print("เลิกงาน")
			} else {
				if t2 > 18.00 {
					print("OT")
				}
			}
		}
	} else if t1 >= 9.30 {
		print("ขาดงาน")
	} else {
		print("สาย")
		if t2 >= 16.00 && t2 <= 18.00 {
			print("เลิกงาน")
		} else {
			if t2 > 18.00 {
				print("OT")
			}
		}
	}
}
