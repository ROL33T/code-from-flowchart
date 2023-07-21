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

func Wait(d time.Duration) {
	time.Sleep(d)
}

const (
	randomStringChars  = "abcdefghijklmnopqrstuvwxyz"
	randomStringLength = 1000
)

func generateRandomString(length int) string {
	var sb strings.Builder
	sb.Grow(length)

	for i := 0; i < length; i++ {
		if i == 0 {
			sb.WriteByte(randomStringChars[rand.Intn(len(randomStringChars)-1)])
		} else {
			sb.WriteByte(randomStringChars[rand.Intn(len(randomStringChars))])
		}
	}

	return sb.String()
}

func findMaxMinCounts(str, sub string) (maxCount, minCount int) {
	startIndex := 0
	maxCount = 0
	minCount = len(str)

	for {
		index := strings.Index(str[startIndex:], sub)
		if index == -1 {
			break
		}

		count := strings.Count(str[startIndex:], sub)
		if count > maxCount {
			maxCount = count
		}

		if count < minCount {
			minCount = count
		}

		startIndex += index + len(sub)
	}

	return maxCount, minCount
}

func findDuplicateCharacters(input string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range input {
		counts[char]++
	}
	return counts
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(randomStringLength)
	printf("ตัวอักษรทั้งหมด:\n[%s] \n", randomString)

	println("กรุณาใส่ String 5 ตัวเท่านั้น |  : ")
	string_What := ""
	if _, err := fmt.Scan(&string_What); err != nil {
		println("ไม่สามารถอ่านค่า Input String ได้")
		Wait(2 * time.Second)
		return
	}

	string_what_lower := strings.ToLower(string_What)

	if len(string_What) == 5 {

		if string_what_lower >= "a" || string_what_lower <= "z" || string_what_lower == "*" {
			countOfAsterisk := strings.Count(string_what_lower, "*")

			if strings.Contains(string_what_lower, "*") {
				for i := 0; i < countOfAsterisk; i++ {
					fmt.Printf("กรอกค่าที่จะแทนที่ * ที่ตำแหน่งที่ %d: ", i+1)
					var replacement string
					if _, err := fmt.Scan(&replacement); err != nil {
						println("ไม่สามารถอ่านค่า Input String ได้")
						return
					}
					if replacement >= "a" && replacement <= "z" {
						string_what_lower = strings.Replace(string_what_lower, "*", replacement, 1)

					} else {
						println("A-Z เท่านั้นพี่ชาย")
						Wait(2 * time.Second)
						return
					}

				}
			}

			string_count := strings.Count(randomString, string_what_lower)

			if string_count > 0 {
				printf("พบ %s %d จำนวน \n", string_what_lower, string_count)

				maxCount, minCount := findMaxMinCounts(randomString, string_what_lower)
				printf("มากที่สุด: %d\n", maxCount)
				printf("น้อยที่สุด: %d\n", minCount)
			} else {
				printf("ไม่พบ %s string นี้ \n", string_what_lower)
			}
		} else {
			println("A-Z และ * เท่านั้นพี่ชาย")
			Wait(2 * time.Second)
			return
		}

	} else {
		println("กรุณากรอกให้ครบ 5 ตัว")
		Wait(2 * time.Second)
		return
	}
	Wait(5 * time.Second)
}
