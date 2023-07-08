package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	var a int
	fmt.Print("กรุณาใส่จำนวนเงินที่ชำระบิลเดือนนี้: ") //ยอดที่ต้องชำระเดือนนี้
	fmt.Scan(&a)

	c := 0  //บัตรกำนัน
	d := "" // ประกาศตัวแปรไว้print

	if a > 15000 {

		var b int
		fmt.Print("กรุณาใส่จำนวนเงินที่ชำระบิลเดือนหน้า: ") //ยอดที่้ต้องชำระเดือนหน้า
		fmt.Scan(&b)

		if b > 10000 {
			c = 1000
			d = "คุณได้รับบัตรกำนัน"
		} else {
			d = "คุณไม่ได้รับบัตรกำนันเนื่องจาก ไม่สามารถตัดบัตรได้เนื่องจากคุณไม่มีเงิน ชำระเดือนหน้า"
		}

	} else {
		d = "คุณไม่ได้รับบัตรกำนันเนื่องจาก ไม่สามารถตัดบัตรได้เนื่องจากคุณไม่มีเงิน ชำระเดือนนี้"
	}

	if c >= 1000 {
		printf("%s %d", d, c)
	} else {
		printf("%s", d)
	}

}
