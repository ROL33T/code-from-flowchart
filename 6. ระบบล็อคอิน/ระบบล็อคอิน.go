package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	user, password := "test", "1234"
	countLogin := 0
	countPass := 0
	expDateLockLogin := time.Now()
	Cooldown := 30 * time.Minute
	lockClient := false
	isLoginSuccess := false

	reader := bufio.NewReader(os.Stdin)

	var userinput, passwordinput string

	for {
		if lockClient {
			diff := expDateLockLogin.Sub(time.Now())
			days := int(diff.Hours()) / 24
			hours := int(diff.Hours()) % 24
			minutes := int(diff.Minutes()) % 60
			seconds := int(diff.Seconds()) % 60

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
			if seconds > 0 {
				durationStr += fmt.Sprintf("%d วินาที", seconds)
			}

			print("ระบบทำการล็อค คุณต้องรอ เวลาถึง", durationStr, "ถึงจะปลดล็อกให้")
			time.Sleep(1 * time.Second)
			continue
		}

		if !isLoginSuccess {
			print("กรุณาใส่ Username: ")
			userinput, _ = reader.ReadString('\n')
			userinput = strings.TrimSpace(userinput)

			print("กรุณาใส่ Password: ")
			passwordinput, _ = reader.ReadString('\n')
			passwordinput = strings.TrimSpace(passwordinput)

			if userinput == user && passwordinput == password {
				print("login ok")
				countLogin = 0
				countPass = 0
				expDateLockLogin = time.Now()
				lockClient = false
				isLoginSuccess = true
				break
			}

			countLogin++
			if countLogin >= 10 {
				expDateLockLogin = time.Now().Add(Cooldown)
				lockClient = true
			} else {
				countPass++
				if countPass >= 5 {
					expDateLockLogin = time.Now().Add(Cooldown)
					lockClient = true
				} else {
					print("รหัสคุณผิดพลาด กรุณาลองใหม่อีกครั้งที่", countPass)
				}
			}
		} else {
			countLogin = 0
			countPass = 0
			lockClient = false
			isLoginSuccess = false
		}
	}

}
