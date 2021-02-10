package main

import "fmt"

var a [3][3]int

func draw() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(a[i][j], " ")
		}
		fmt.Println()
	}
}

func checker() int {

	if a[0][0] == a[0][1] && a[0][0] == a[0][2] {
		return a[0][0]
	} else if a[0][0] == a[1][1] && a[0][0] == a[2][2] {
		return a[0][0]
	} else if a[0][0] == a[1][0] && a[0][0] == a[2][0] {
		return a[0][0]
	} else if a[1][0] == a[1][1] && a[1][0] == a[1][2] {
		return a[1][0]
	} else if a[2][0] == a[2][1] && a[2][0] == a[2][2] {
		return a[2][0]
	} else if a[0][1] == a[1][1] && a[0][1] == a[2][1] {
		return a[0][1]
	} else if a[0][2] == a[1][2] && a[2][2] == a[0][2] {
		return a[0][2]
	} else if a[0][2] == a[1][1] && a[0][2] == a[2][0] {
		return a[0][2]
	}

	return 1

}
func main() {

	var x, y, flag, count int
	for {
		if flag == 0 {
			fmt.Println("First Players Turn:")
			fmt.Scan(&x)
			fmt.Scan(&y)
			if (x > 3 || x < 0) || (y > 3 || y < 0) {
				fmt.Println("Please enter valid coordinates")
				continue
			}
			if a[x][y] == 0 {
				count++
				flag++
				a[x][y] = 1
				draw()
			}
		} else if flag == 1 {
			fmt.Println("Second  Players Turn:")
			fmt.Scan(&x)
			fmt.Scan(&y)
			if (x > 3 || x < 0) || (y > 3 || y < 0) {
				fmt.Println("Please enter valid coordinates")
				continue
			}
			if a[x][y] == 0 {
				count++
				flag++
				a[x][y] = 1
				draw()
			}
		}

		if count > 8 {
			fmt.Println("It's a Tie Guys!")
			break
		}
		var v int
		v = checker()
		if v == 1 {
			fmt.Println("First Player Won!Congratualtions")
			break
		} else if v == 2 {
			fmt.Println("Second Player Won!Congratulations")
			break
		}

	}

}
