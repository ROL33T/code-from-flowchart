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

	/*
			เขียนโค๊ดแสดงรูป
		    x
		   oxo
		  aoxoa
		 waoxoaw
		qwaoxoawq
		 waoxoaw
		  aoxoa
		   oxo
		    x

		โดยใช้ Loop-for เท่านั้น

	*/
	lines := []string{
		"     x",
		"    oxo",
		"   aoxoa",
		"  waoxoaw",
		" qwaoxoawq",
		"  waoxoaw",
		"   aoxoa",
		"    oxo",
		"     x",
	}

	for _, line := range lines {
		println(line)
	}
}
