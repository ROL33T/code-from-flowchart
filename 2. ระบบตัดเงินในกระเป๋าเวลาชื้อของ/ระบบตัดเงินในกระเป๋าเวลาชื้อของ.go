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
	var money int
	fmt.Print("กรุณาใส่จำนวนเงินในกระเป๋า: ")
	fmt.Scan(&money)

	var pay int
	fmt.Print("กรุณาใส่ราคาสิ่งของ: ")
	fmt.Scan(&pay)

	if money >= pay {
		money = money - pay
		printf("ระบบทำการตัดเงินเรียบร้อยแล้ว ตอนนี้ ยอดเงินเหลือ %d", money)
	} else {
		printf("ระบบไม่สามารถตัดเงินคุณได้เนื่องจาก เงินของคุณไม่พอ ต้องการอีก %d", pay-money)
	}
}
