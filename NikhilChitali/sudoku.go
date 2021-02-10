package main

import "fmt"

var board = [9][9]int {
	{5,3,0,0,7,0,0,0,0},
	{6,0,0,1,9,5,0,0,0},
	{0,9,8,0,0,0,0,6,0},
	{8,0,0,0,6,0,0,0,3},
	{4,0,0,8,0,3,0,0,1},
	{7,0,0,0,2,0,0,0,6},
	{0,6,0,0,0,0,2,8,0},
	{0,0,0,4,1,9,0,0,5},
	{0,0,0,0,8,0,0,7,9}}

func main() {
	
	var i,j int
	for i = 0; i < 9; i++ {
		for j = 0; j < 9; j++ {
			fmt.Printf("%d", board[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Solution")

	fmt.Println()
	fmt.Println()	
	solve(0,0)
}

func solve(row,col int) {
	if row == 9 {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				fmt.Printf("%d", board[i][j])
			}
			fmt.Println()
 		}
		return
	}
	
	newrow := 0
	newcol := 0

	if col == 8 {
		newrow = row + 1
		newcol = 0
	} else {	
		newrow = row
		newcol = col + 1
	}
	 

	if board[row][col] != 0 {
		solve(newrow, newcol)
	} else {
		for put := 1; put <= 9; put++ {
			if isvalidposition(row, col, put) == true {
				board[row][col] = put
				solve(newrow, newcol)
				board[row][col] = 0			
			}
		}
	}

	return
}

func isvalidposition (row,col,val int) bool {
	for j := 0; j < 9; j++ {
		if board[row][j] == val {
			return false
		}
	}

	for i:= 0; i < 9; i++ {
		if board[i][col] == val {
			return false		
		}
	}

	var blockrow = row/3 * 3
	var blockcol = col/3 * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i + blockrow][j + blockcol] == val {
				return false
			}
		} 
	}

	return true
 }
