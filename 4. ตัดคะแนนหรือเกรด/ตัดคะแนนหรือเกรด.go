package main

import "fmt"

func print(a ...interface{}) {
	fmt.Println(a...)
}

func Scan(a ...interface{}) {
	fmt.Scan(a...)
}

func printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func calculateGrade(score int) string {
	var grade string

	if score > 90 {
		grade = "A"
	} else if score >= 85 && score <= 90 {
		grade = "B+"
	} else if score >= 80 && score <= 85 {
		grade = "B"
	} else if score >= 75 && score <= 80 {
		grade = "C+"
	} else if score >= 70 && score <= 75 {
		grade = "C"
	} else if score >= 60 && score <= 70 {
		grade = "D+"
	} else if score >= 55 && score <= 60 {
		grade = "D"
	} else if score >= 0 && score <= 55 {
		grade = "F"
	}

	return grade
}

func main() {
	var a int
	print("กรุณาใส่คะแนนงาน: ") //คะแนนงาน
	Scan(&a)

	var b int
	print("กรุณาใส่คะแนนเช็คชื่อ: ") //คะแนนเช็คชื่อ
	Scan(&b)

	c := 0             //คะแนนพิเศษ
	d := 0             //คะแนน Final
	f := a + b + c + d // รวม คะแนนทั้งหมด
	g := "F"           //เอาไว้ประกาศ ตัวแปร เกรด

	if b < 5 {
		g = "F"
	} else {
		if a >= 40 {
			if b >= 10 {
				d = 5
				f = a + b + c + d
				g = calculateGrade(f)
			} else {
				d = 0
				f = a + b + c + d
				g = calculateGrade(f)
			}
		} else {
			d = 0
			f = a + b + c + d
			g = calculateGrade(f)
		}
	}
	printf("คุณได้คะแนน %d เกรด: %s\n", f, g)
}
