package main

import (
	"fmt"
	"strings"
	"time"
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
	//ระบบแจ้งเตือนการปิดปรับปรุงระบบ กำหนดเวลาปิดปรับปรุงระบบคือ 22.50 - 02.00 ของทุกวัน โดยหากป้อนข้อมูลเวลาเข้าไปหากอยู่ในช่วงเวลาที่กำหนดให้แจ้งว่าระบบปิดปรับปรุง
	//*กำหนดให้ใช้ package time*
	println("กรุณาใส่ เวลา เช่น 16:00 |  : ")
	time_input := ""
	Scan(&time_input)

	timeParts := strings.Split(time_input, ":")

	timeInput := timeParts[0] + ":" + timeParts[1]

	time_Current, err := time.Parse("15:04", timeInput)

	if err != nil {
		println("เกิดข้อผิดพลาดในการแปลงเวลา:", err)
		return
	}

	if time_Current.Hour() >= 22 && time_Current.Minute() >= 50 || time_Current.Hour() >= 23 || time_Current.Hour() <= 2 && time_Current.Minute() <= 0 {
		println("เชิฟเวอร์ปิดปรับปรุง จนถึง ตี 02:00")

	} else {
		println("เชิฟเวอร์ไม่ได้ปิด")
	}
}
