package main

import (
	"fmt"
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

// class GiftBox

type GiftBox struct {
	Level int
}

type GiftManager struct {
	GiftBoxes []*GiftBox
}

func (gm *GiftManager) AddGiftBox(g *GiftBox) {
	gm.GiftBoxes = append(gm.GiftBoxes, g)
}

func (gm *GiftManager) LogF(format string, a ...interface{}) {
	printf("[Log] "+format, a...)
}

func (g *GiftBox) Open() {
	printf("[Log] Give gift box level %d\n", g.Level)
}

func main() {
	var levelCurrent int
	print("กรุณาใส่เวเวลที่คุณต้องการ: ") // เวเวล ล่าสุด
	Scan(&levelCurrent)

	var countDayRegister int
	print("กรุณาใส่วันที่สมัคร เช่น 1 วัน หรือ 7 วัน: ") //วันที่สมัคร
	Scan(&countDayRegister)

	var eventDateStart int                                     // เวลาล่าสุด ของกิจกรรม
	print("กรุณาใส่วันกิจกรรมของคุณ เช่น 1 วัน หรือ 14 วัน: ") //วันที่กิจกรรม
	Scan(&eventDateStart)
	eventDateEnd := 14 // เวลาล่าสุด + เพิ่ม 7 วันหลังจากเริ่มกิจกรรม

	gm := GiftManager{}
	countGiftCache := 0 // จำนวนครั้งที่รับของ
	X := 0              //คำนวน เวเวล

	if countDayRegister <= 7 {
		gm.LogF("%s", "Welcome gift box 1 ea\n")
		if eventDateStart > eventDateEnd {
			gm.LogF("%s", "เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			countCurrent := levelCurrent / 5
			for i := 0; i < countCurrent; i++ {
				if levelCurrent >= 5 {
					X = X + 5
					g := &GiftBox{Level: X}
					g.Open()
					gm.AddGiftBox(g)
					countGiftCache++
					levelCurrent -= 5
				} else {
					gm.LogF("%s", "คุณไม่ได้รับของ")
				}
			}
		}
	} else {
		gm.LogF("%s", "ไม่ได้รับของเนื่องจากคุณสมัครเกิน 7 วันแล้ว")
		if eventDateStart > eventDateEnd {
			gm.LogF("%s", "เกินเวลาที่กิจกรรมกำหนดไว้")
		} else {
			countCurrent := levelCurrent / 5
			for i := 0; i <= countCurrent; i++ {
				if levelCurrent >= 5 {
					X = X + 5
					g := &GiftBox{Level: X}
					g.Open()
					gm.AddGiftBox(g)
					countGiftCache++
					levelCurrent -= 5
				} else {
					gm.LogF("%s", "คุณไม่ได้รับของ")
				}
			}
		}
	}

	gm.LogF("คุณได้รับ Give gift box level ทั้งหมด %d กล่อง\n", countGiftCache)
}
