package main

import "fmt"

var board [3][3]string

func draw() {
	fmt.Println(" ", board[0][0], "  |  ", board[0][1], "  |  ", board[0][2])
	fmt.Println("--------------")
	fmt.Println(" ", board[1][0], "  |  ", board[1][1], "  |  ", board[1][2])
	fmt.Println("--------------")
	fmt.Println(" ", board[2][0], "  |  ", board[2][1], "  |  ", board[2][2])
}

func check(x, y, player int) bool {
	if player == 1 && board[x][y] == "" {
		board[x][y] = "X"
	} else if player == 2 && board[x][y] == "" {
		board[x][y] = "O"
	} else {
		return false
	}

	var flagx, flagy, flagd1, flagd2 int = 0, 0, 0, 0
	if x == y {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == j && board[x][y] == board[i][j] {
					flagd1++
				}
			}
		}
	}

	if x+y == 2 {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i+j == 2 && board[x][y] == board[i][j] {
					flagd2++
				}
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == x && board[i][j] == board[x][y] {
				flagx++
			}

			if j == y && board[i][j] == board[x][y] {
				flagy++
			}
		}
	}
	if flagx == 3 {
		return true
	}
	if flagy == 3 {
		return true
	}
	if flagd1 == 3 {
		return true
	}
	if flagd2 == 3 {
		return true
	}
	return false
}

func main() {
	draw()
	var first, second string
	fmt.Print("Enter name of Player1=")
	fmt.Scan(&first)
	fmt.Print("Enter name of Player2=")
	fmt.Scan(&second)

	var x, y, i int

	var turn int = 1

	for i = 0; i < 9; i++ {
		if turn == 1 {
			fmt.Println(first, " enter x and y coordinates:")
			fmt.Scan(&x, &y)
			if check(x, y, 1) {
				fmt.Println(first, " Won!!")
				break
			}
			turn = 2
			draw()
		} else {
			fmt.Println(second, " enter x and y coordinates:")
			fmt.Scan(&x, &y)
			if check(x, y, 2) {
				fmt.Println(second, " Won!!")
				break
			}
			turn = 1
			draw()
		}
	}
	if i == 9 {
		fmt.Println("It is a Draw :(")
		draw()
	}
}
