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

type NumberEven struct {
	Value int64
}
type NumberOdd struct {
	Value int64
}

func main() {
	a := 0 //นับเลขคู่
	b := 0 //นับเลขคี่
	var numbers_even []NumberEven
	var numbers_odd []NumberOdd

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			a++
			//println("คู่", i)
			numbers_even = append(numbers_even, NumberEven{Value: int64(i)})

		} else {
			b++
			//println("คี่", i)
			numbers_odd = append(numbers_odd, NumberOdd{Value: int64(i)})
		}

	}

	println(numbers_even)
	println(numbers_odd)
}
