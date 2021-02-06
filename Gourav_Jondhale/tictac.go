package main

import "fmt"

var game [3][3] int

func draw(){
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Printf("%v ",game[i][j])
		}
		fmt.Println()
	}
}


func result() int{

	if(game[0][0]==game[0][1] && game[0][0]==game[0][2]){
		return game[0][0]
	} else if(game[0][0]==game[1][1] && game[0][0]==game[2][2]){
		return game[0][0]
	} else if(game[0][0]==game[1][0] && game[0][0]==game[2][0]){
		return game[0][0]
	} else if(game[1][0]==game[1][1] && game[1][0]==game[1][2]){
		return game[1][0]
	} else if(game[2][0]==game[2][1] && game[2][0]==game[2][2]){
		return game[2][0]
	} else if(game[0][1]==game[1][1] && game[0][1]==game[2][1]){
		return game[0][1]
	} else if(game[0][2]==game[1][2] && game[2][2]==game[0][2]){
		return game[0][2]
	} else if(game[0][2]==game[1][1] && game[0][2]==game[2][0]){
		return game[0][2]
	}
	return 5
}

func main(){
	
	var x,y,flg,cnt int
	draw()
	for {
		
		// fmt.Printf("%v, %v",x,y)
		if(flg==0){
			fmt.Println("Player 1 turn: ")
			fmt.Scan(&x)
			fmt.Scan(&y)
			if ((x>3 || x<0) || (y>3 || y<0)){
				fmt.Println("Please enter valid coordinates")
				continue
			}
			fmt.Println()
			if(game[x][y]==0){
				game[x][y]=1
				cnt++
				flg=1
			draw()
			}
		} else {
			fmt.Println("Player 2 turn: ")
			fmt.Scan(&x)
			fmt.Scan(&y)
			if ((x>3 || x<0) || (y>3 || y<0)){
				fmt.Println("Please enter valid coordinates")
				continue
			}
			fmt.Println()
			if(game[x][y]==0){
				game[x][y]=2
				cnt++
				flg=0
				draw()
			}
		}
	
		if(cnt>9){
			fmt.Println("It's a tie!! ")
			break
		}
	
		var w int
		w=result()
		if(w==5){
			continue
		}
		if(w==1){
		fmt.Printf("Winner is Player %v: ",w)
		break
		} else if(w==2){
			fmt.Printf("Winner is Player %v: ",w)
			break
		} 
	}

	
}
