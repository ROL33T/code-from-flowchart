package main

import (
	"fmt"
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
	// แสดงผลข้อมูลปฏิทินเดือนที่ 6-12 โดยให้ print ในรูปแบบของปฏิทิน

	const calendarFormat = `
Sun Mon Tue Wed Thu Fri Sat 
%s`

	// วันเริ่มต้นและวันสิ้นสุดของเดือนที่ต้องการแสดงผล
	startMonth := time.June
	endMonth := time.December

	// วนลูปเพื่อแสดงผลปฏิทิน
	for month := startMonth; month <= endMonth; month++ {
		// ประมวลผลเดือนและปีปัจจุบัน
		now := time.Now()
		year := now.Year()

		// กำหนดวันแรกของเดือน
		firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

		// คำนวณจำนวนวันในเดือน
		nextMonth := month + 1
		if nextMonth > endMonth {
			nextMonth = time.January
			year++
		}
		lastDayOfMonth := time.Date(year, nextMonth, 0, 0, 0, 0, 0, time.UTC)
		numDays := lastDayOfMonth.Day()

		// พิมพ์หัวข้อเดือน
		monthHeader := month.String()
		printf("\n----------%s----------\n", monthHeader)

		// พิมพ์หัวของวัน
		printf(calendarFormat, "")

		// พิมพ์วันของสัปดาห์แรก
		firstWeekday := firstDayOfMonth.Weekday()
		for i := 0; i < int(firstWeekday); i++ {
			printf("    ")
		}

		// พิมพ์วันที่
		for day := 1; day <= numDays; day++ {
			if day == time.Now().Day() && month == time.Now().Month() {
				printf("")
			} else {
				printf("%3d ", day)
			}

			// เมื่อเป็นวันอาทิตย์ ขึ้นบรรทัดใหม่
			if (day+int(firstWeekday))%7 == 0 {
				println()
			}
		}

		println()
	}
}
