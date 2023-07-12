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
	//เขียนการทำงานระบบแสดงวันหากได้รับข้อมูลวันที่ Ex. Input -> 01/07/2023 Output -> Saturday
	println("กรุณาใส่ วัน:เดือน:ปี : ")
	time_input := ""
	Scan(&time_input)

	timeParts := strings.Split(time_input, "/")

	timeInput := timeParts[0] + "/" + timeParts[1] + "/" + timeParts[2]

	day_Current, err := time.Parse("02/01/2006", timeInput)
	if err != nil {
		println("เกิดข้อผิดพลาดในการแปลงวัน:", err)
		return
	}

	println(day_Current.Weekday())
}
