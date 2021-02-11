//This program allows a user to play sudoku
package main

import (
	"fmt"
	"strconv"
)

// var matrix = [9][9]string {
// 	{"_","_","_","2","6","_","7","_","1"},
// 	{"6","8","_","_","7","_","_","9","_"},
// 	{"1","9","_","_","_","4","5","_","_"},
// 	{"8","2","_","1","_","_","_","4","_"},
// 	{"_","_","4","6","_","2","9","_","_"},
// 	{"_","5","_","_","_","3","_","2","8"},
// 	{"_","_","9","3","_","_","_","7","4"},
// 	{"_","4","_","_","5","_","_","3","6"},
// 	{"7","_","3","_","1","8","_","_","_"}}

var matrix = [9][9]string{
	{"4", "3", "5", "2", "6", "9", "7", "8", "1"},
	{"6", "8", "2", "5", "7", "1", "4", "9", "3"},
	{"1", "9", "7", "8", "3", "4", "5", "6", "2"},
	{"8", "2", "6", "1", "9", "5", "3", "4", "7"},
	{"3", "7", "4", "6", "8", "2", "9", "1", "5"},
	{"9", "5", "1", "7", "4", "3", "6", "2", "_"},
	{"5", "1", "9", "3", "2", "6", "8", "7", "_"},
	{"2", "4", "8", "9", "5", "7", "1", "3", "_"},
	{"7", "6", "3", "4", "1", "8", "2", "5", "9"}}

//Counts the number of blank spaces remaining
func countBlanks()int{
	i:=0
	for _,rowv := range matrix{
		for _,colv := range rowv {
			if(colv == "_"){
				i++
			}
		}		
	}
	return i
}

var blankSpaces int = countBlanks()

func checkDuplicates(row int, col int, guess string) bool {
	//Check for same row
	for _, v:= range matrix[row] {
		if(v == guess){
			// fmt.Println("This integer already exists in row", row+1)
			return false;
		}
	}

	//Check for same col
	for _,v := range matrix {
		if(v[col] == guess){
			// fmt.Println("This integer already exists in column", col+1)
			return false;
		}
	}

	//Check within same box
	var rbox int = (row/3)*3 
	var cbox int = (col/3)*3 

	for i:= rbox; i < rbox+3; i++ {
		for j:= cbox; j < cbox+3; j++ {
			if(matrix[i][j] == guess){
				// fmt.Println("This integer already exists in its box")
				return false
			}
		}
	}

	//Return true only if all checks have passed
	return true;
}


func main(){
	for  {
		//Prints the placeholder
		fmt.Println(blankSpaces, "spaces left to fill")
		fmt.Println("   1 2 3     4 5 6     7 8 9")
		fmt.Println("  ---------------------------")
		for i:=0; i<9; i++ {
			fmt.Println(i+1,matrix[i][:3],"|",matrix[i][3:6],"|", matrix[i][6:])
			if((i+1)%3 == 0){
				fmt.Println("  ---------------------------")
			}
		}
		//Loop will exit when blank spaces are 0
		if(blankSpaces == 0){
			break;
		}
		//Take a row input
		rowInp:
		var row int
		fmt.Println("Row:");
		fmt.Scanln(&row)
		//Validate the row input
		if(!(0 < row && row < 10)){
			fmt.Println("Please input a valid row between 1 - 9")
			goto rowInp
		}

		//Take a col input
		colInp:
		var col int
		fmt.Println("Col:");
		fmt.Scanln(&col)
		//Validate the col input
		if(!(0 < col && col < 10)){
			fmt.Println("Please input a valid column between 1 - 9")
			goto colInp
		}

		//Take a guess input
		guessInp:
		var guess string
		fmt.Println("Guess:(Type q to quit)")
		fmt.Scanln(&guess)
		//Handle quit
		if(guess == "q"){
			fmt.Println("Thanks for playing!")
			return;
		}
		//Handle blank input
		if(guess == "_"){
			matrix[row-1][col-1] = "_"
			blankSpaces++
			continue
		}
		//Handle guess
		//Validate guess input
		guessInt, _ := strconv.Atoi(guess)
		if(!(0 < guessInt && guessInt < 10)){
			fmt.Println("Please input a valid guess between 1 - 9 to make a guess, _ to clear a cell or q to quit")
			goto guessInp
		}
		//Validate the guess within the sudoku matrix
		valid := checkDuplicates(row-1, col-1, guess)
		//Update matrix for correct guess
		if(valid){
			fmt.Println("Your guess is valid!")
			matrix[row-1][col-1] = guess
			blankSpaces--
		}
	}

	//If the for loop is exited, the sudoku must have been solved
	fmt.Println("Congratulations! You've solved the sudoku!")
	return
}