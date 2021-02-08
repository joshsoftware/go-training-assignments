// Tic-Tac-Toe considering 3X3 board

package main

import "fmt"

func check_win(matrix [][]string, value string) bool {
	var won = false
	// diagonal match
	if (matrix[0][0] == value && matrix[1][1] == value && matrix[2][2] == value) ||
		(matrix[0][2] == value && matrix[1][1] == value && matrix[2][0] == value) {
		return true
	}
	// vertical or horizontal match
	for i := 0; i < 3; i++ {
		if (matrix[i][0] == value && matrix[i][1] == value && matrix[i][2] == value) ||
			(matrix[0][i] == value && matrix[1][i] == value && matrix[2][i] == value) {
			won = true
			break
		}
	}
	return won
}

func switch_player(chance_of string) string {
	if chance_of == "X" {
		return "O"
	} else {
		return "X"
	}
}

func place_valid_and_available(matrix [][]string, x, y int) bool {
	if (x > -1 && x < 3) && (y > -1 && y < 3) && matrix[x][y] == "-" {
		return true
	} else {
		return false
	}
}

func draw(matrix [][]string) {
	// print play board
	for i := 0; i < 3; i++ {
		fmt.Println(matrix[i])
	}
}

func main() {

	var chance_of string = "X"
	var chances = 9
	var x, y int

	// Create Play board
	play_board := make([][]string, 3)
	for i := 0; i < 3; i++ {
		play_board[i] = make([]string, 3)
	}

	// assign default value as "-"
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			play_board[i][j] = "-"
		}
	}

	// main code
	for chances > -1 {
		if chances == 0 {
			fmt.Println("Draw!!")
			break
		} else {
			fmt.Println("Player ", chance_of,
				"'s turn \nEnter space seperated x y co-ordinates")
			fmt.Scan(&x, &y)
			if place_valid_and_available(play_board, x, y) {
				play_board[x][y] = chance_of
				chances -= 1
				draw(play_board)
				if check_win(play_board, chance_of) {
					fmt.Println(chance_of, "Won!!")
					break
				}
				chance_of = switch_player(chance_of)
			} else {
				fmt.Println("Place is either invalid or already taken")
			}
		}
	}
}

// ==================================================================
// OUTPUT

// akansha@akansha-Latitude-3490:~/golang/assignments$ go run main.go
// Player  X 's turn
// Enter space seperated x y co-ordinates
// 0 0
// [X - -]
// [- - -]
// [- - -]
// Player  O 's turn
// Enter space seperated x y co-ordinates
// 2 0
// [X - -]
// [- - -]
// [O - -]
// Player  X 's turn
// Enter space seperated x y co-ordinates
// 0 1
// [X X -]
// [- - -]
// [O - -]
// Player  O 's turn
// Enter space seperated x y co-ordinates
// 0 2
// [X X O]
// [- - -]
// [O - -]
// Player  X 's turn
// Enter space seperated x y co-ordinates
// 1 1
// [X X O]
// [- X -]
// [O - -]
// Player  O 's turn
// Enter space seperated x y co-ordinates
// 2 1
// [X X O]
// [- X -]
// [O O -]
// Player  X 's turn
// Enter space seperated x y co-ordinates
// 2 2
// [X X O]
// [- X -]
// [O O X]
// X Won!!
// akansha@akansha-Latitude-3490:~/golang/assignments$
