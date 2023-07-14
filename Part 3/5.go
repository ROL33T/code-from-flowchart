package main

import (
	"fmt"
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

func Scanln(a ...interface{}) {
	fmt.Scanln(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func main() {
	var input string
	println("string แล้วนำมาเรียงตัวอักษรจากมา:")
	Scanln(&input)

	words := strings.Split(input, ",")

	sort.Strings(words)

	sortedString := strings.Join(words, ",")

	println(sortedString)
}
