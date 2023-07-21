package main

import (
	"fmt"
	"math/rand"
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

func generateRandomString(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteByte(randomStringChars[rand.Intn(len(randomStringChars))])
	}

	return sb.String()
}

const (
	randomStringChars  = "abcdefghijklmnopqrstuvwxyz"
	randomStringLength = 1000
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(randomStringLength)
	println("กรุณาใส่ String 5 ตัวเท่านั้น |  : ")
	string_What := ""
	if _, err := fmt.Scan(&string_What); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		time.Sleep(2 * time.Second)
		return
	}

	string_lower := strings.ToLower(string_What)

	if len(string_lower) == 5 {
		string_count := strings.Count(randomString, string_lower)

		startIndex := 0

		for {
			index := strings.Index(randomString[startIndex:], string_lower)
			if index == -1 {
				break
			}

			printf("พบคำว่า %s ที่ตำแหน่งที่: %d\n", string_lower, startIndex+index)

			startIndex += index + len(string_lower)
		}

		if string_count > 0 {
			printf("พบ %s จำนวนทั้งหมด: %d\n", string_lower, string_count)
		} else {
			printf("ไม่พบ %s", string_lower)
		}

	} else {
		println("กรุณากรอกให้ครบ 5 ตัว")
		time.Sleep(2 * time.Second)
		return
	}
	time.Sleep(5 * time.Second)
}
