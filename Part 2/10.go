package main

import "fmt"

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

func main() {
	rows := 9
	space := rows / 2
	num := 1

	words := []string{"", "x", "o", "a", "w", "q"}

	for i := 1; i <= rows; i++ {

		for j := 1; j <= space; j++ {
			printf(" ")
		}

		count := num/2 + 1

		for k := 1; k <= num; k++ {
			printf("%s", words[count])
			if k <= num/2 {
				count--
			} else {
				count++
			}
		}

		println()

		if i <= rows/2 {
			space--
			num += 2
		} else {
			space++
			num -= 2
		}

	}

}
