package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func main() {

	/*
		เขียนเครื่องคิดเลข + - * / ที่สามารถคำนวณได้หลายรอบและหากได้รับคำว่า reset ระบบจะรีการคำนวณเป็น 0 และถ้าได้รับคำว่า exit คือการออกจากระบบ
	*/
	scanner := bufio.NewScanner(os.Stdin)

	result := 0
	exit := false

	for !exit {
		println("เครื่องคิดเลข")
		println("ใช้คำสั่ง 'clear หรือ cls' เพื่อล้างหน้าต่าง TERMINAL")
		println("ใช้คำสั่ง 'reset' เพื่อรีเซ็ตค่า")
		println("ใช้คำสั่ง 'exit' เพื่อออกจากโปรแกรม")
		println("เช่น 40 + 40 หรือ 20 - 40 หรือ 10 * 40 หรือ 1 / 40")
		print("ป้อนตัวดำเนินการและตัวเลข หรือ คำสั่ง: ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())

		if input == "reset" {
			result = 0
			println("ค่าถูกรีเซ็ตเป็น 0")
			continue
		} else if input == "clear" || input == "cls" {
			for i := 1; i <= 15; i++ {
				println("")
			}

			println("ล้างหน้า TERMINAL เรียบร้อยแล้ว")
			continue

		} else if input == "exit" {
			exit = true
			println("ออกจากโปรแกรม")
			break
		}

		parts := strings.Split(input, " ")
		if len(parts) != 3 {
			println("รูปแบบไม่ถูกต้อง กรุณาลองอีกครั้ง")
			continue
		}
		numberone, err := strconv.Atoi(parts[0])
		if err != nil {
			println("ตัวเลขไม่ถูกต้อง กรุณาลองอีกครั้ง")
			continue
		}

		operator := parts[1]

		numbertwo, err := strconv.Atoi(parts[2])
		if err != nil {
			println("ตัวเลขไม่ถูกต้อง กรุณาลองอีกครั้ง")
			continue
		}

		switch operator {
		case "+":
			result_num := numberone + numbertwo
			result += result_num
		case "-":
			result_num := numberone - numbertwo
			result -= result_num
		case "*":
			result_num := numberone * numbertwo
			if result > 0 {
				result *= result_num
			} else {
				result = result_num
			}
		case "/":
			if numberone != 0 || numbertwo != 0 {
				result_num := numberone / numbertwo
				result /= result_num
			} else {
				println("ไม่สามารถหารด้วยศูนย์ได้")
			}
		default:
			println("ตัวดำเนินการไม่ถูกต้อง กรุณาลองอีกครั้ง")
		}

		println("ผลลัพธ์:", result)
	}
}
