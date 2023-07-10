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
	print("กรุณาใส่เวลาเข้างาน: ") //เวลาเข้างาน
	time_input := ""
	Scan(&time_input)

	timeParts := strings.Split(time_input, ":")

	timeInput := timeParts[0] + ":" + timeParts[1]
	time_Current, err := time.Parse("15:04", timeInput)

	if err != nil {
		fmt.Println("เกิดข้อผิดพลาดในการแปลงเวลา:", err)
		return
	}

	JOIN_WORK := false

	if time_Current.Hour() <= 8 && time_Current.Minute() <= 59 || time_Current.Hour() == 9 && time_Current.Minute() == 0 {
		println("เข้างาน")
		JOIN_WORK = true
	} else if time_Current.Hour() >= 9 && time_Current.Minute() >= 30 {
		println("ขาดงาน")
		JOIN_WORK = false
	} else {
		println("สาย")
		JOIN_WORK = true
	}

	if JOIN_WORK {
		print("กรุณาใส่เวลาออกงาน: ") //เวลาออกงาน
		time_out_input := ""
		Scan(&time_out_input)
		timeParts := strings.Split(time_out_input, ":")

		timeInput := timeParts[0] + ":" + timeParts[1]
		time_Current_Out, err := time.Parse("15:04", timeInput)
		if err != nil {
			fmt.Println("เกิดข้อผิดพลาดในการแปลงเวลา:", err)
			return
		}

		if time_Current_Out.Hour() <= 15 && time_Current_Out.Minute() <= 59 {
			println("ขาดงาน")
		} else {
			JOIN_WORK = true
		}
		
		if time_Current_Out.Hour() >= 16 && time_Current_Out.Minute() <= 59 && time_Current_Out.Hour() <= 17 && time_Current_Out.Minute() <= 59 {
			println("เลิกงาน")
		} else {
			if time_Current_Out.Hour() >= 18 {
				println("OT")
			}
		}
	} 
}
