package main

import (
	"fmt"
)

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
	var money int
	print("กรุณาใส่จำนวนเงินในกระเป๋า: ")
	Scan(&money)

	var pay int
	print("กรุณาใส่ราคาสิ่งของ: ")
	Scan(&pay)

	if money >= pay {
		money = money - pay
		printf("ระบบทำการตัดเงินเรียบร้อยแล้ว ตอนนี้ ยอดเงินเหลือ %d", money)
	} else {
		printf("ระบบไม่สามารถตัดเงินคุณได้เนื่องจาก เงินของคุณไม่พอ ต้องการอีก %d", pay-money)
	}
}
