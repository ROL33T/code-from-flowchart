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
	money := 100 //เงินในกระเป๋า

	pay := 500 //ยอดที่ต้องชำระ

	if money >= pay {
		printf("ระบบทำการตัดเงินเรียบร้อยแล้ว ตอนนี้ ยอดเงินเหลือ %d", money-pay)
	} else {
		printf("ระบบไม่สามารถตัดเงินคุณได้เนื่องจาก เงินของคุณไม่พอ ต้องการอีก %d", pay-money)
	}
}
