package main

import "fmt"

func main() {
	defer fmt.Println("It was a good game.")
	matrix := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	break_continue := true
	var first, second int
	var is_valid bool
	var filled_position int
	var player_no int
	var center_cube_covered_by int
	var is_center_cube_covered bool

	for (break_continue) {
		is_valid = true
		if filled_position % 2 == 0 {

			player_no = 1
			fmt.Println("Player 1 Chance")
		} else {
			player_no = 2
			fmt.Println("Player 2 Chance")
		}

		for (is_valid) {
			fmt.Println("please enter the first index")
			fmt.Scanln(&first)
		
			fmt.Println("please enter the second index")
			fmt.Scanln(&second)

			if(matrix[first][second] == 0) {
				matrix[first][second] = player_no
				is_valid = false
				filled_position += 1
			} else {
				fmt.Println("The position is preoccupied. Please choose an empty position.")
				printMatrix(matrix)
			}
		}

		if (first == 1 && second == 1) {
			center_cube_covered_by = player_no
			is_center_cube_covered = true
		}

		if (matrix[first][0] == player_no && matrix[first][1] == player_no && matrix[first][2] == player_no) {
			fmt.Println("player %d win.", player_no)
			break
		}

		if (matrix[0][second] == player_no && matrix[1][second] == player_no && matrix[2][second] == player_no) {
			fmt.Println("player %d win.", player_no)
			break
		}

		if (is_center_cube_covered && center_cube_covered_by == player_no) {
			if((matrix[0][0] == player_no && matrix[2][2] == player_no) || (matrix[0][2] == player_no && matrix[2][0] == player_no)) {
				fmt.Println("player %d win.", player_no)
				break
			}
		}

		if(filled_position == 9) {
			fmt.Println("Good Game! it's a draw.")
			break;
		}
		printMatrix(matrix)
	}
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for _, value := range matrix {
		fmt.Println(value)
	}
}
