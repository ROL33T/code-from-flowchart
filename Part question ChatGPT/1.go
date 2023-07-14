package main

import (
	"fmt"
)

func FindMissingNumber(arr []int) int {
	n := len(arr) + 1 // จำนวนเต็มที่ควรจะมีในอาเรย์
	totalSum := (n * (n + 1)) / 2
	arrSum := 0

	for _, num := range arr {
		arrSum += num
	}

	missingNumber := totalSum - arrSum
	return missingNumber
}
func main() {
	arr := []int{1, 2, 4, 6, 3, 7, 8}
	missingNumber := FindMissingNumber(arr)
	fmt.Println(missingNumber) // ผลลัพธ์: 5 (จำนวนเต็มที่ขาดหายไปในอาเรย์คือ 5)

}
