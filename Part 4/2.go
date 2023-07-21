package main

import (
	"fmt"
	"math/rand"
	"sort"
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

func sortString(s string) string {

	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })

	return string(runes)
}

func countCharacters(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	return counts
}

const (
	randomStringChars  = "abcdefghijklmnopqrstuvwxyz"
	randomStringLength = 1000
)

func main() {

	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(randomStringLength)

	printf("ตัวอักษรทั้งหมด:\n[%s] \n", randomString)

	find_string_one := ""

	println("กรุณาใส่ Input Two |  : ")

	string_lower_find_one := strings.ToLower(find_string_one)

	if _, err := fmt.Scan(&string_lower_find_one); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		time.Sleep(2 * time.Second)
		return
	}

	if len(string_lower_find_one) == 2 {

		string_count := strings.Count(randomString, string_lower_find_one)

		startIndex := 0

		for {
			index := strings.Index(randomString[startIndex:], string_lower_find_one)
			if index == -1 {
				break
			}

			printf("พบคำว่า %s ที่ตำแหน่งที่: %d\n", string_lower_find_one, startIndex+index)

			startIndex += index + len(string_lower_find_one)
		}

		if string_count > 0 {
			printf("พบ %s จำนวนทั้งหมด: %d\n", string_lower_find_one, string_count)
			time.Sleep(2 * time.Second)
			return
		} else {
			printf("ไม่พบ %s", string_lower_find_one)
			time.Sleep(2 * time.Second)
			return
		}

	} else {
		println("กรุณากรอกให้ครบ 2 ตัว")
		time.Sleep(2 * time.Second)
		return
	}

	time.Sleep(5 * time.Second)
}
