package main

import (
	"fmt"
)

//Initializes the matrix with a row
var matrix = [9][9]string {
	{"2","1","3","4","5","6","7","8","9"}}
//Initializes the possibilities matrix which will store the values of
//possibilities for each cell in a slice.
var possibilitiesMatrix=[9][9][]string{}
//This will store the slice of last valid possibilities for each cell
//which will be referred in case of regression
var lastValidPossibilitiesMatrix = [9][9][]string{}
var digits = [9]string{"1","2","3","4","5","6","7","8","9"}
//Regressed indicates the main function which matrix to use for recursion
var regressed bool = false

//Checks the duplicates for a guess for a particular cell
func checkDuplicates(row int, col int, guess string) bool {
	//Check for same row
	for _, v:= range matrix[row] {
		if(v == guess){
			return false;
		}
	}

	//Check for same col
	for _,v := range matrix {
		if(v[col] == guess){
			return false;
		}
	}

	//Check within same box
	var rbox int = (row/3)*3 
	var cbox int = (col/3)*3 

	for i:= rbox; i < rbox+3; i++ {
		for j:= cbox; j < cbox+3; j++ {
			if(matrix[i][j] == guess){
				return false
			}
		}
	}
	//Return true only if all checks have passed
	return true;
}

//Appends to the slice only if it is a unique value
func appendIfUnique(arr *[]string, str string){
	present:=false
	for _,v:=range *arr{
		if(str == v){
			present = true
			break
		}
	}
	if(!present){
		*arr = append(*arr, str)
	}
}

//Takes the current matrix, generates the possibilities for every empty cell
//and updates the possibilitiesMatrix. If any cell returns 0 possibilities, 
//it returns false. Else, it returns true.
func checkPossibilites()bool{
	var possibilities = [9][9][]string{}
	for i,row:=range matrix{
		for j,cell:=range row {
			if cell == ""{
				for _,dig:= range digits{
					valid:= checkDuplicates(i, j, dig)
					if(valid){
						appendIfUnique(&possibilities[i][j], dig)
					}
				}
				if len(possibilities[i][j]) == 0{
					return false
				}
			}
		}
	}
	possibilitiesMatrix = possibilities
	return true
}

//Prints the matrix
func printMatrix(){
	for _,v:=range matrix{
		fmt.Println(v)
	}
}


//In case checkPossibilities fails at a cell, this resets the matrix to the 
//previous cell and uses the lastValidPossibilitiesMatrix to move to try the 
//next option
func regress(i, j int){
	matrix[i][j]="";
	if(j==0){
		matrix[i-1][8]=""
		lastValidPossibilitiesMatrix[i-1][8]=lastValidPossibilitiesMatrix[i-1][8][1:]
	} else {
		matrix[i][j-1]=""
		lastValidPossibilitiesMatrix[i][j-1]=lastValidPossibilitiesMatrix[i][j-1][1:]
	}
	regressed = true
}

func main(){
	//Sets the initial possibilities matrix
	checkPossibilites()

	regress:
	for i, row:= range matrix {
		for j, cell:= range row {
			//Only executes the code if the cell is empty
			if cell == ""{
				//If the regressed flag is true, it uses the lastValidPossibilitiesMatrix
				//to evaluate new possibilities.
				if(regressed){
					//If the length of the lastValidPossibilitiesMatrix is 0,
					//It will regress one step further and skip the rest of the process
					if(len(lastValidPossibilitiesMatrix[i][j]) == 0){
						regress(i,j)
						goto regress
					}

					for _,poss:=range lastValidPossibilitiesMatrix[i][j]{
						//Assigns the current possibility to the matrix cell
						matrix[i][j] = poss
						//Checks possibilities for new value of the cell
						valid:=checkPossibilites()
						//If a valid possibility is found, moves to the next cell
						if(valid){
							regressed = false
							break
						}
						//If no options are valid, regresses to previous cell
						if(len(lastValidPossibilitiesMatrix[i][j]) == 1){
							regress(i,j)
							goto regress
						}
					}
				} else {
					//If regressed is false, the possibilitiesMatrix is used for 
					//evaluating possibilities
					for _,poss:= range possibilitiesMatrix[i][j]{
						//This stores the current possibilites for the cell into lastValidPossibilitiesMatrix.
						//If a possibility is valid, this valid will be saved in the lastValidMatrix
						lastValidPossibilitiesMatrix[i][j] = possibilitiesMatrix[i][j]
						//Assigns the current possibility to the matrix cell
						matrix[i][j] = poss
						//Checks possibilities for new value of the cell
						valid:=checkPossibilites()
						//If a valid possibility is found, moves to the next cell
						if(valid){
							break;
						} 
						//If no options are valid, regresses to previous cell
						if(len(possibilitiesMatrix[i][j]) == 1){
							regress(i,j)
							goto regress
						}
					}
				}
			}
		}
	}	
	
	fmt.Printf("You have a sudoku matrix!\n")
	for _,v:=range matrix{
		fmt.Println(v)
	}
}