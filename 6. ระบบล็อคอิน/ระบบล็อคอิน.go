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

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

var userinput, passwordinput string

func main() {

	user, password := "test", "1234"

	countLogin := 0
	countPass := 0
	expDateLockLogin := time.Now()
	lockClient := false

	reader := bufio.NewReader(os.Stdin)

	for {
		print("Enter username input: ")
		userinput, _ = reader.ReadString('\r')
		userinput = strings.TrimSpace(userinput)

		print("Enter password input: ")
		passwordinput, _ = reader.ReadString('\r')
		passwordinput = strings.TrimSpace(passwordinput)

		if expDateLockLogin.After(time.Now()) {
			print("reset")
			countLogin = 0
			countPass = 0
			expDateLockLogin = time.Time{}
			lockClient = false
		} else {
			if lockClient {
				print("lock user")
				break
			} else {
				if userinput == user {
					if passwordinput == password {
						print("login ok")
						countLogin = 0
						countPass = 0
						expDateLockLogin = time.Now()
						lockClient = false
						break // Exit the loop on successful login
					} else {
						countPass++
						if countPass >= 5 {
							expDateLockLogin = expDateLockLogin.Add(-30 * time.Minute)
							lockClient = true
						} else {
							print("รหัสคุณผิดพลาด กรุณาลองใหม่อีกครั้งที่", countPass)
						}
					}
				} else {
					countLogin++
					if countLogin >= 10 {
						expDateLockLogin = expDateLockLogin.Add(-30 * time.Minute)
						lockClient = true
					} else {
						print("UserName คุณผิดพลาด กรุณาลองใหม่อีกครั้ง", countLogin)
					}
				}
			}
		}
	}

}
