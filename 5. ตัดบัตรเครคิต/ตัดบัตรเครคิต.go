package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	a := 0 //ยอดที่ต้องชำระเดือนนี้
	b := 12000//ยอดที่้ต้องชำระเดือนหน้า
	c := 0 //บัตรกำนัน
	d := "" // ประกาศตัวแปรไว้print

	if a > 15000 {
		if b > 10000 {
			c = 1000
			d = "คุณได้รับบัตรกำนัน"
		}
	} else {
		d = "คุณไม่ได้รับบัตรกำนันเนื่องจาก ไม่สามารถตัดบัตรได้เนื่องจากคุณไม่มีเงิน"
	}

	if c >= 1000 {
		printf("%s %d", d, c)
	} else {
		printf("%s", d)
	}

}
