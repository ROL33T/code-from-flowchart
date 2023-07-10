package main

import (
	"fmt"
	"strconv"
	"strings"
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

	if len(timeParts) != 2 {
		println("รูปแบบเวลาไม่ถูกต้อง")
		return
	}

	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		println("ชั่วโมงไม่ถูกต้อง")
		return
	}

	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		println("นาทีไม่ถูกต้อง")
		return
	}

	JOIN_WORK := false

	if hour <= 9 && minute <= 00 {
		println("เข้างาน")
		JOIN_WORK = true
	} else if hour >= 9 && minute >= 30 {
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

		if len(timeParts) != 2 {
			println("รูปแบบเวลาไม่ถูกต้อง")
			return
		}

		hour_out, err := strconv.Atoi(timeParts[0])
		if err != nil {
			println("ชั่วโมงไม่ถูกต้อง")
			return
		}

		minute_out, err := strconv.Atoi(timeParts[1])
		if err != nil {
			println("นาทีไม่ถูกต้อง")
			return
		}

		if hour_out <= 15 && minute_out <= 59 {
			println("ขาดงาน")
		} else {
			JOIN_WORK = true
		}

		if hour_out >= 16 && minute_out >= 00 && hour_out <= 18 && minute_out <= 00 {
			println("เลิกงาน")
		} else {
			if hour_out >= 18 && minute_out > 0 {
				println("OT")
			}
		}

	}
}
