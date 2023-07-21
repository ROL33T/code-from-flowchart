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

	sortedString := sortString(randomString)

	find_string_one := ""

	println("กรุณาใส่ Input Two |  : ")

	string_lower_find_one := strings.ToLower(find_string_one)

	if _, err := fmt.Scan(&string_lower_find_one); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		return
	}

	if len(string_lower_find_one) == 1 {
		find_string_two := ""

		println("กรุณาใส่ Input Two |  : ")

		string_lower_find_two := strings.ToLower(find_string_two)

		if _, err := fmt.Scan(&string_lower_find_two); err != nil {
			println("ไม่สามารถอ่านค่า Input Two ได้")
			return
		}

		if len(string_lower_find_two) == 1 {
			counts := countCharacters(sortedString)

			for char, count := range counts {
				if string(char) == string_lower_find_one {
					printf("%c: %d\n", char, count)
				}

				if string(char) == string_lower_find_two {
					printf("%c: %d\n", char, count)
				}
			}
		}
	}
	time.Sleep(5 * time.Second)
}
