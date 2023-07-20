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

func countCharactersAndDigits(input string) (int, int, int) {
	charCount := 0
	digitCount := 0
	symbolCount := 0

	for _, char := range input {
		if unicode.IsLetter(char) {
			charCount++
		} else if unicode.IsDigit(char) {
			digitCount++
		} else if unicode.IsSymbol(char) || char == '-' || char == '*' || char == '/' || char == '+' {
			symbolCount++
		}
	}

	return charCount, digitCount, symbolCount
}

func main() {

	println("กรุณาใส่ข้อมูล: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	characters, digits, symbol := countCharactersAndDigits(input)

	printf("LETTERS %d DIGITS %d SYMBOL %d\n", characters, digits, symbol)
}
