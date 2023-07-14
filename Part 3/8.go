package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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

func processText(text string) string {
	words := strings.Fields(text)

	uniqueWords := make([]string, 0)

	seen := make(map[string]bool)

	for _, word := range words {
		if !seen[word] {
			seen[word] = true
			uniqueWords = append(uniqueWords, word)
		}
	}

	sort.Strings(uniqueWords)

	result := strings.Join(uniqueWords, " ")

	return result
}

func main() {
	println("กรุณาใส่ข้อความ: ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')

	userInput = strings.TrimSpace(userInput)

	output := processText(userInput)

	println("ผลลัพธ์:", output)
}
