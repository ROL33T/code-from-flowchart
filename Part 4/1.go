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

	sortedString := sortString(randomString)

	counts := countCharacters(sortedString)

	for char, count := range counts {
		printf("%c: %d\n", char, count)
	}

	Wait(5 * time.Second)

}
