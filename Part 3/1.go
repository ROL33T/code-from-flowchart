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

func ShowCalc(number int) {
	max_val := 12
	for i := 1; i <= max_val; i++ {
		printf("%d x %d : %d", number, i, number*i)
		println()
	}
}

func main() {
	input_number := 0
	print("กรุณาใส่เลขเพื่อแสดงผลสูตรคูณ เช่น 1 :")
	Scan(&input_number)

	ShowCalc(input_number)

}
