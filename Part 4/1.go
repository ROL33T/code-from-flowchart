package main

import (
	"fmt"
	"math/rand"
	"regexp"
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

func countCharacters_A(s, substr string) int {
	count := 0
	if strings.Contains(substr, "*") {
		escapedSubstr := strings.ReplaceAll(substr, "*", "\\*")
		re := regexp.MustCompile(escapedSubstr)
		matches := re.FindAllStringIndex(s, -1)
		count = len(matches)
	} else {
		for i := 0; i <= len(s)-len(substr); i++ {
			if s[i:i+len(substr)] == substr {
				count++
			}
		}
	}
	return count
}

const (
	randomStringChars  = "abcdefghijklmnopqrstuvwxyz"
	randomStringLength = 1000
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomString := generateRandomString(randomStringLength)

	printf("ตัวอักษรทั้งหมด:\n[%s]\n", randomString)

	sortedString := sortString(randomString)

	counts := countCharacters(sortedString)
	Wait(1 * time.Second)
	//------1
	for char, count := range counts {
		printf("%c: %d\n", char, count)
	}
	//------ 2
	find_string_one := ""

	println("กรุณาใส่ Input String 2 ตัวเท่านั้น |  : ")

	string_lower_find_one := strings.ToLower(find_string_one)

	if _, err := fmt.Scan(&string_lower_find_one); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		Wait(2 * time.Second)
		return
	}

	if string_lower_find_one >= "a" || string_lower_find_one <= "z" {
		if len(string_lower_find_one) == 2 {
			string_count := countCharacters_A(randomString, string_lower_find_one)
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
				Wait(2 * time.Second)

			} else {
				printf("ไม่พบ %s", string_lower_find_one)
				Wait(2 * time.Second)
				return
			}
		} else {
			println("กรุณากรอกให้ครบ 2 ตัว")
			Wait(2 * time.Second)
			return
		}
	} else {
		println("A-Z เท่านั้นพี่ชาย")
		Wait(2 * time.Second)
		return
	}

	//------3

	println("กรุณาใส่ String 5 ตัวเท่านั้น |  : ")

	string_What_A := ""
	if _, err := fmt.Scan(&string_What_A); err != nil {
		println("ไม่สามารถอ่านค่า Input One ได้")
		Wait(2 * time.Second)
		return
	}

	string_lower_A := strings.ToLower(string_What_A)

	if string_lower_A >= "a" || string_lower_A <= "z" {
		if len(string_lower_A) == 5 {
			string_count := countCharacters_A(randomString, string_lower_A)
			startIndex := 0
			for {
				index := strings.Index(randomString[startIndex:], string_lower_A)
				if index == -1 {
					break
				}

				printf("พบคำว่า %s ที่ตำแหน่งที่: %d\n", string_lower_A, startIndex+index)

				startIndex += index + len(string_lower_A)
			}
			if string_count > 0 {
				printf("พบ %s จำนวนทั้งหมด: %d\n", string_lower_A, string_count)
			} else {
				printf("ไม่พบ %s", string_lower_A)
			}
		} else {
			println("กรุณากรอกให้ครบ 5 ตัว")
			Wait(2 * time.Second)
			return
		}
	} else {
		println("A-Z เท่านั้นพี่ชาย")
		Wait(2 * time.Second)
		return
	}

	//-----4

	ARRAY_MAX_TO_MIN := map[string]int{}

	for {
		println("กรุณาใส่ String 5 * ตัวเท่านั้น | end จบ  : ")
		string_What := ""
		if _, err := fmt.Scan(&string_What); err != nil {
			println("ไม่สามารถอ่านค่า Input String ได้")
			Wait(2 * time.Second)
			return
		}

		string_what_lower := strings.ToLower(string_What)

		if string_what_lower >= "0" && string_what_lower <= "9" {
			println("A-Z และ * เท่านั้นพี่ชาย")
			Wait(2 * time.Second)
			return
		}

		if string_what_lower == "end" {
			println("จบการทำงาน")
			Wait(2 * time.Second)
			return
		}

		if len(string_what_lower) == 5 {
			if string_what_lower >= "a" || string_what_lower <= "z" || string_what_lower == "*" {
				countOfAsterisk := countCharacters_A(string_what_lower, "*")
				if countOfAsterisk == 1 || countOfAsterisk == 0 {
					if strings.Contains(string_what_lower, "*") {
						for i := 0; i < countOfAsterisk; i++ {
							for char := 'a'; char <= 'z'; char++ {
								string_what_lower := strings.Replace(string_what_lower, "*", string(char), 1) //replace = แทนที่
								string_count := countCharacters_A(randomString, string_what_lower)

								if string_count > 0 {
									ARRAY_MAX_TO_MIN[string_what_lower]++

									printf("พบ %s %d จำนวน\n", string_what_lower, string_count)

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
								}
							}
						}
					} else {
						string_count := countCharacters_A(randomString, string_what_lower)

						if string_count > 0 {

							ARRAY_MAX_TO_MIN[string_what_lower]++

							printf("พบ %s %d จำนวน\n", string_what_lower, string_count)

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
							printf("ไม่พบ %s string นี้ \n", string_what_lower)
						}
					}
				} else {
					println("ไม่รองรับ * มากกว่า 1 ตัว:")
					Wait(1 * time.Second)
					return
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
	}
	Wait(5 * time.Second)
}
