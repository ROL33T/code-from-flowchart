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
	user := "test" // username กำหนดเอาไว้
	password := "test" // password กำหนดเอาไว้
	userinput := "test1" // input username
	passwordinput := "test1" // input password
	countLogin := 0 // จำนวนครั้งที่ล็อคอินผิด
	countPass := 0 // จำนวนครั้งที่ใส่พาสผิด
	expDateLockLogin := time.Now() //เวลาล่าสุด
	lockClient := false // ล็อค User หรือ ล็อค ไอพี หรือ ไม่ true = ล็อค false = ไม่ล็อค

	if expDateLockLogin.After(time.Now()) {
		print("reset")
		countLogin = 0
		countPass = 0
		expDateLockLogin = time.Time{}
		lockClient = false
	} else {
		if lockClient {
			print("lock user")
		} else {
			if user == userinput {
				if password == passwordinput {
					print("login ok")
					countLogin = 0
					countPass = 0
					expDateLockLogin = time.Now()
					lockClient = false
				} else {
					countPass++
					if countPass >= 5 {
						expDateLockLogin = expDateLockLogin.Add(time.Duration(-30) * time.Minute)
						lockClient = true
					} else {
						print("รหัสคุณผิดพลาด กรุณาลองใหม่อีกครั้งที่ ",countPass)
					}
				}
			} else {
				countLogin++
				if countLogin >= 10 {
					expDateLockLogin = expDateLockLogin.Add(time.Duration(-30) * time.Minute)
					lockClient = true
				} else {
					print("UserName คุณผิดพลาด กรุณาลองใหม่อีกครั้ง ",countLogin)
				}
			}
		}
	}
}
