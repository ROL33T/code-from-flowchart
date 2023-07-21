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

const (
	randomStringChars  = "abcdefghijklmnopqrstuvwxyz*"
	randomStringLength = 1000
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(randomStringLength)

	string_What := ""
	if _, err := fmt.Scan(&string_What); err != nil {
		println("ไม่สามารถอ่านค่า Input String ได้")
		return
	}

	string_what_lower := strings.ToLower(string_What)

	if len(string_What) == 5 {

		countOfAsterisk := strings.Count(string_what_lower, "*")

		if strings.Contains(string_what_lower, "*") {
			string_what_lower = strings.Replace(string_what_lower, "*", string(randomStringChars[rand.Intn(26)]), countOfAsterisk)
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
		println("กรุณากรอกให้ครบ 5 ตัว")
		time.Sleep(2 * time.Second)
		return
	}
	time.Sleep(5 * time.Second)
}
