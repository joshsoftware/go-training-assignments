//This program solves the sudoku puzzle by itself.
package main

import(
	"fmt"
)

var matrix = [9][9]string {
	{"_","_","_","2","6","_","7","_","1"},
	{"6","8","_","_","7","_","_","9","_"},
	{"1","9","_","_","_","4","5","_","_"},
	{"8","2","_","1","_","_","_","4","_"},
	{"_","_","4","6","_","2","9","_","_"},
	{"_","5","_","_","_","3","_","2","8"},
	{"_","_","9","3","_","_","_","7","4"},
	{"_","4","_","_","5","_","_","3","6"},
	{"7","_","3","_","1","8","_","_","_"}}

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

var digits = []string{"1","2","3","4","5","6","7","8","9"}
	
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

func updatePossibilitiesMatrix()[9][9][]string{
	var possibilitiesMatrix = [9][9][]string{}
	for i,row:= range possibilitiesMatrix {
		for j,_:= range row {
			if(matrix[i][j] == "_"){
				for _,dig:=range digits{	
					valid := checkDuplicates(i, j, dig)
					if(valid){
						appendIfUnique(&possibilitiesMatrix[i][j], dig)
					}
				}
			}
			
		}
	}

	return possibilitiesMatrix
}

func updateSolutionMatrix(possibilitiesMatrix [9][9][]string){
	for i,row:=range possibilitiesMatrix{
		for j,cell:=range row{
			if(len(cell) == 0){
				continue
			}
			if(len(cell)==1){
				matrix[i][j]=cell[0]
				blankSpaces--
				possibilitiesMatrix[i][j]=[]string{}
				continue
			}
		}
	}

	fmt.Println("UPDATING SOLUTION")
	for _,row := range matrix{
		fmt.Println(row)
	} 
}

func main(){
	var possibilitiesMatrix = [9][9][]string{}
	i:=0;
	for ;i<10;i++{
		possibilitiesMatrix = updatePossibilitiesMatrix()
		updateSolutionMatrix(possibilitiesMatrix)
		if(blankSpaces == 0){
			break;
		}
	}

	if(blankSpaces == 0){
		fmt.Printf("The Sudoku has been solved in %v iterations!\n", i+1)
		for _,v:=range matrix{
			fmt.Println(v)
		}
	}
}