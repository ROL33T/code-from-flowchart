package main

import "fmt"

func GetKey(key int) (int, int) {
	int_key := 0
	int_val := 0

	if key < 0 || key == 0 || key > 9 {
		int_key = 200
		int_val = 200
	}

	if key == 1 {
		int_key = 0
		int_val = 0
	}
	if key == 2 {
		int_key = 0
		int_val = 1
	}
	if key == 3 {
		int_key = 0
		int_val = 2
	}
	if key == 4 {
		int_key = 1
		int_val = 0
	}
	if key == 5 {
		int_key = 1
		int_val = 1
	}
	if key == 6 {
		int_key = 1
		int_val = 2
	}
	if key == 7 {
		int_key = 2
		int_val = 0
	}
	if key == 8 {
		int_key = 2
		int_val = 1
	}
	if key == 9 {
		int_key = 2
		int_val = 2
	}

	return int_key, int_val
}

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

func showBoard(board [][]string) {
	for _, row := range board {
		println(row)
	}
}

func getPlayerMove(turn string) int {
	var move int
	fmt.Printf("Enter your Turn %s move (1-9): ", turn)
	fmt.Scan(&move)
	if move >= 1 && move <= 9 {
		return move
	} else {
		return 0
	}

}

func checkWin(board [][]string, symbol string) bool {
	// Check vertical lines
	for col := 0; col < 3; col++ {
		if board[0][col] == symbol && board[1][col] == symbol && board[2][col] == symbol {
			return true
		}
	}

	// Check horizontal lines
	for row := 0; row < 3; row++ {
		if board[row][0] == symbol && board[row][1] == symbol && board[row][2] == symbol {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == symbol && board[1][1] == symbol && board[2][2] == symbol {
		return true
	}
	if board[0][2] == symbol && board[1][1] == symbol && board[2][0] == symbol {
		return true
	}

	return false
}

func isMoveAvailable(board [][]string, move int) bool {
	row, col := getRowAndCol(move)
	if board[row][col] == "_" {
		return true
	}
	return false
}

func getRowAndCol(move int) (int, int) {
	row := (move - 1) / 3
	col := (move - 1) % 3
	return row, col
}

func isBoardFull(board [][]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == "_" {
				return false
			}
		}
	}
	return true
}

func resetBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = "_"
		}
	}
}

func main() {
	Turn := "X"
	var board = [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	for {
		if checkWin(board, "X") {
			println("ผู้เล่น X ชนะ!")
			break
		} else if checkWin(board, "O") {
			println("ผู้เล่น O ชนะ!")
			break
		} else if isBoardFull(board) {
			println("เสมอ เล่นใหม่!")
			resetBoard(board)
		} else {
			move := getPlayerMove(Turn)

			if move == 0 {
				println("ช่องนี้ไม่มีอยู่จริง!")
			} else {

				intKey, intVal := GetKey(move)

				if isMoveAvailable(board, move) {
					if Turn == "X" {
						board[intKey][intVal] = "X"
						Turn = "O"
					} else {
						board[intKey][intVal] = "O"
						Turn = "X"
					}
				} else {
					println("ช่องนี้ไม่ว่างกรุณาเลือกช่องใหม่")
				}
			}
		}
		showBoard(board)
	}
}
