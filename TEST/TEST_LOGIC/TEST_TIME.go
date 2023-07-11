package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	YYYYMMDD = "2006-01-02"

	HHMM = "00:00"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now.Format(YYYYMMDD))
	time_input := ""
	fmt.Scan(&time_input)
	TIME_FORMAT := fmt.Sprintf("%v %v", now.Format(YYYYMMDD), time_input)
	fmt.Println(TIME_FORMAT)

	year_now := time.Now().Year()
	month_now := time.Now().Month()
	day_now := time.Now().Day()
	second_now := time.Now().Second()

	timeParts := strings.Split(time_input, ":")

	if len(timeParts) != 2 {
		fmt.Println("รูปแบบเวลาไม่ถูกต้อง")
		return
	}

	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		fmt.Println("ชั่วโมงไม่ถูกต้อง")
		return
	}

	minute, err := strconv.Atoi(timeParts[1])
	if err != nil {
		fmt.Println("นาทีไม่ถูกต้อง")
		return
	}

	t := time.Date(year_now, month_now, day_now, hour, minute, second_now, 0, time.UTC)

	fmt.Println(t.Unix())
}
