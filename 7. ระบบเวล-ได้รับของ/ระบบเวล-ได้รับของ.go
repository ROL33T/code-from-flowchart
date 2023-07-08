package main

import (
	"fmt"
	"time"
)

func print(a ...interface{}) {
	fmt.Println(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	levelCurrent := 21                                     // เวเวล ล่าสุด
	countDayRegister := 1                                  //วันที่สมัคร
	eventDateStart := time.Now()                           // เวลาล่าสุด ของกิจกรรม
	eventDateEnd := eventDateStart.Add(7 * 24 * time.Hour) // เวลาล่าสุด + เพื่ม 7 วันหลังจากเรื่ม กิจกรรม
	countGiftCache := 0                                    //จำนวนครั้งที่รับของ
	countCurrent := 0

	if countDayRegister <= 7 {
		print("welcome gift box 1ea")
		if eventDateStart.After(eventDateEnd) {
			print("เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			countCurrent = int(levelCurrent / 5)
			for i := 0; i < int(countCurrent); i++ {
				if levelCurrent >= 5 {
					levelCurrent -= 5
					countGiftCache++
					print("gift box level 1 ea")
				} else {
					print("คุณไม่ได้รับของ")
				}
			}

		}
	} else {
		print("ไม่ได้รับของเนื่องจากคุณสมัครเกิน 7 วันแล้ว")
		if eventDateStart.After(eventDateEnd) {
			print("เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			countCurrent = int(levelCurrent / 5)
			for i := 0; i < countCurrent; i++ {
				if levelCurrent >= 5 {
					levelCurrent -= 5
					countGiftCache++
					print("gift box level 1 ea")
				} else {
					print("คุณไม่ได้รับของ")
				}
			}

		}
	}

	printf("คุณได้รับ giftbox level ทั้งหมด %d กล่อง", countGiftCache)
}
