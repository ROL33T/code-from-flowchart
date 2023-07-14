package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
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

func countCharactersAndDigits(input string) (int, int) {
	charCount := 0
	digitCount := 0

	for _, char := range input {
		if unicode.IsLetter(char) {
			charCount++
		} else if unicode.IsDigit(char) {
			digitCount++
		}
	}

	return charCount, digitCount
}

func main() {

	println("กรุณาใส่ข้อมูล: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	characters, digits := countCharactersAndDigits(input)

	printf("LETTERS %d DIGITS %d\n", characters, digits)
}
