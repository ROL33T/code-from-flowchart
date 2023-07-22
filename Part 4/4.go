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

func getInputString() string {
	fmt.Print("กรุณาใส่ String 5 ตัวเท่านั้น | end จบ: ")
	var inputString string
	fmt.Scan(&inputString)
	return inputString
}

func validateInputString(inputString string) bool {
	if len(inputString) != 5 {
		return false
	}

	for _, char := range inputString {
		if (char < 'a' || char > 'z') && char != '*' {
			return false
		}
	}

	return true
}

func replaceAsterisk(inputString string) string {
	countOfAsterisk := strings.Count(inputString, "*")

	for i := 0; i < countOfAsterisk; i++ {
		fmt.Printf("กรอกค่าที่จะแทนที่ * ที่ตำแหน่งที่ %d: ", i+1)
		var replacement string
		fmt.Scan(&replacement)

		if replacement >= "a" && replacement <= "z" {
			inputString = strings.Replace(inputString, "*", replacement, 1)
		} else {
			fmt.Println("A-Z เท่านั้นพี่ชาย")
			time.Sleep(2 * time.Second)
			return ""
		}
	}

	return inputString
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ARRAY_MAX_TO_MIN := map[string]int{}
	randomString := generateRandomString(randomStringLength)
	fmt.Printf("ตัวอักษรทั้งหมด:\n[%s]\n", randomString)

	for {
		inputString := getInputString()
		stringWhatLower := strings.ToLower(inputString)

		if stringWhatLower == "end" {
			println("จบการทำงาน")
			Wait(2 * time.Second)
			return
		}

		if !validateInputString(stringWhatLower) {
			println("กรุณากรอกให้ครบ 5 ตัว และเป็น A-Z หรือ * เท่านั้นพี่ชาย")
			Wait(2 * time.Second)
			return
		}

		stringWhatLower = replaceAsterisk(stringWhatLower)

		if stringWhatLower == "" {
			Wait(2 * time.Second)
			return
		}

		stringCount := strings.Count(randomString, stringWhatLower)

		if stringCount > 0 {
			ARRAY_MAX_TO_MIN[stringWhatLower]++
			printf("พบ %s %d จำนวน\n", stringWhatLower, stringCount)

			// หา key ใน map ที่เรียงลำดับตามค่า value จากมากไปน้อย
			sortedKeys := make([]string, 0, len(ARRAY_MAX_TO_MIN))
			for key := range ARRAY_MAX_TO_MIN {
				sortedKeys = append(sortedKeys, key)
			}

			sort.SliceStable(sortedKeys, func(i, j int) bool {
				return ARRAY_MAX_TO_MIN[sortedKeys[i]] > ARRAY_MAX_TO_MIN[sortedKeys[j]]
			})

			println("เรียงลำดับตามค่าจากมากไปน้อย:")
			for _, key := range sortedKeys {
				value := ARRAY_MAX_TO_MIN[key]
				printf("%s: %d\n", key, value)
			}
		} else {
			printf("ไม่พบ %s string นี้\n", stringWhatLower)
		}
	}
	Wait(5 * time.Second)
}
