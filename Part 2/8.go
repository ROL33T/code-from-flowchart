package main

import (
	"fmt"
)

func println(a ...interface{}) {
	fmt.Println(a...)
}

func print(a ...interface{}) {
	fmt.Print(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	//t := []string{“mon”,”wed”,”sat”,”sun”,”tue”,”fri”,”thu”} จาก array เก็บวัน ที่กำหนดจงเรียงค่าวันเป็น sun,mon,tue,wed,thu,fri,sat แล้วแสดง array นั้นออกมา
	t := []string{"mon", "wed", "sat", "sun", "tue", "fri", "thu"}

	t[0] = "sun"
	t[1] = "mon"
	t[2] = "tue"
	t[3] = "wed"
	t[4] = "thu"
	t[5] = "fri"
	t[6] = "sat"

	println(t)
}
