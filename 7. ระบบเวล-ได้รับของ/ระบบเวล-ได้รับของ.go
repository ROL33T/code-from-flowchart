package main

import (
	"fmt"
	"time"
)

func print(a ...interface{}) {
	fmt.Println(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

type GiftBox struct {
	Level int
}

func (g *GiftBox) Open() {
	fmt.Printf("Give gift box level %d\n", g.Level)
}

type GiftManager struct {
	GiftBoxes []*GiftBox
}

func (gm *GiftManager) AddGiftBox(g *GiftBox) {
	gm.GiftBoxes = append(gm.GiftBoxes, g)
}

func main() {
	var levelCurrent int
	fmt.Print("กรุณาใส่เวเวลที่คุณต้องการ: ") // เวเวล ล่าสุด
	fmt.Scan(&levelCurrent)

	var countDayRegister int
	fmt.Print("กรุณาใส่วันที่สมัคร เช่น 1 วัน หรือ 14 วัน: ") //วันที่สมัคร
	fmt.Scan(&countDayRegister)

	eventDateStart := time.Now()                           // เวลาล่าสุด ของกิจกรรม
	eventDateEnd := eventDateStart.Add(7 * 24 * time.Hour) // เวลาล่าสุด + เพิ่ม 7 วันหลังจากเริ่มกิจกรรม

	gm := GiftManager{}
	countGiftCache := 0 // จำนวนครั้งที่รับของ

	if countDayRegister <= 7 {
		print("Welcome gift box 1 ea")
		if eventDateStart.After(eventDateEnd) {
			print("เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			countCurrent := levelCurrent / 5
			for i := 0; i < countCurrent; i++ {
				if levelCurrent >= 5 {
					levelCurrent -= 5
					g := &GiftBox{Level: 1}
					g.Open()
					gm.AddGiftBox(g)
					countGiftCache++
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
			countCurrent := levelCurrent / 5
			for i := 0; i < countCurrent; i++ {
				if levelCurrent >= 5 {
					levelCurrent -= 5
					g := &GiftBox{Level: 1}
					g.Open()
					gm.AddGiftBox(g)
					countGiftCache++
				} else {
					print("คุณไม่ได้รับของ")
				}
			}
		}
	}

	printf("คุณได้รับ gift box level ทั้งหมด %d กล่อง\n", countGiftCache)
}
