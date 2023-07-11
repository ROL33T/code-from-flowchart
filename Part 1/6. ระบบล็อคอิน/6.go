package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	userinput := ""
	passwordinput := ""
	user, password := "test", "1234"
	countLogin := 0
	countPass := 0
	LoginNow := time.Now()
	expDateLockLogin := time.Now()
	Cooldown := 30 * time.Minute
	lockClient := false
	isLoginSuccess := false

	reader := bufio.NewReader(os.Stdin)

	for {
		LoginNow = time.Now()

		if LoginNow.After(expDateLockLogin) {
			if lockClient {
				println("ระบบได้ทำการปลดล็อค User ของคุณเรียบร้อยแล้ว")
				countLogin = 0
				countPass = 0
				expDateLockLogin = time.Now()
				lockClient = false
			} else {
				if !isLoginSuccess {
					print("กรุณาใส่ Username: ")
					userinput, _ = reader.ReadString('\r')
					userinput = strings.TrimSpace(userinput)

					print("กรุณาใส่ Password: ")
					passwordinput, _ = reader.ReadString('\r')
					passwordinput = strings.TrimSpace(passwordinput)

					if userinput == user {
						if passwordinput == password {
							println("login ok")
							countLogin = 0
							countPass = 0
							expDateLockLogin = time.Now()
							lockClient = false
							isLoginSuccess = true
							// รีเช็ทตัวแปร ทั้งหมด
							break // ออกลูป
						} else {
							countPass++
							if countPass >= 5 {
								expDateLockLogin = time.Now().Add(Cooldown)
								lockClient = true
							} else {
								println("รหัสคุณผิดพลาด กรุณาลองใหม่อีกครั้งที่", countPass)
							}
						}
					} else {
						countLogin++
						if countLogin >= 10 {
							expDateLockLogin = time.Now().Add(Cooldown)
							lockClient = true
						} else {
							println("UserName คุณผิดพลาด กรุณาลองใหม่อีกครั้ง", countLogin)
						}
					}
				} else {
					// รีเช็ทตัวแปร ทั้งหมด
					countLogin = 0
					countPass = 0
					lockClient = false
					isLoginSuccess = false
				}
			}
		} else {
			if lockClient {
				diff := expDateLockLogin.Sub(LoginNow)
				days := int(diff.Hours()) / 24
				hours := int(diff.Hours()) % 24
				minutes := int(diff.Minutes()) % 60
				seconds := int(math.Round(diff.Seconds())) % 60

				durationStr := ""

				if days > 0 {
					durationStr += fmt.Sprintf("%d วัน ", days)
				}
				if hours > 0 {
					durationStr += fmt.Sprintf("%d ชั่วโมง ", hours)
				}
				if minutes > 0 {
					durationStr += fmt.Sprintf("%d นาที ", minutes)
				}
				if seconds >= 0 {
					durationStr += fmt.Sprintf("%d วินาที", seconds)
				}
				println("ระบบทำการล็อค คุณต้องรอ เวลาถึง ", durationStr, " ถึงจะปลดล็อกให้")
			}
			time.Sleep(1000 * time.Millisecond)
		}

	}

}
