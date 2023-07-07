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
	levelCurrent := 5 // เวเวล ล่าสุด
	countDayRegister := 1 //วันที่สมัคร
	eventDateStart := time.Now() // เวลาล่าสุด ของกิจกรรม
	eventDateEnd := eventDateStart.Add(7 * 24 * time.Hour) // เวลาล่าสุด + เพื่ม 7 วันหลังจากเรื่ม กิจกรรม
	countGiftCache := 0 //จำนวนครั้งที่รับของ

	if countDayRegister <= 7 {
		print("welcome gift box 1ea")
		if eventDateStart.After(eventDateEnd) {
			print("เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			if levelCurrent%5 == 0 {
				countGiftCache++
				print("gift box level 1 ea")
			} else {
				print("คุณไม่ได้รับของ")
			}
		}
	} else {
		print("ไม่ได้รับของเนื่องจากคุณสมัครเกิน 7 วันแล้ว")
		if eventDateStart.After(eventDateEnd) {
			print("เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			if levelCurrent%5 == 0 {
				countGiftCache++
				print("gift box level 1 ea")
			} else {
				print("คุณไม่ได้รับของ")
			}
		}
	}
}
