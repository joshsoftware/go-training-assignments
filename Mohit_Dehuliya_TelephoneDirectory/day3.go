package main

import (
	"bufio"
	"fmt"
	"os"
)

type details struct {
	name    string
	address string
	mobile  uint64
	pincode int
}

var m = make(map[int]details)

func adddetails() {
	var id int
	var d details
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Unique ID Assoicated with dis")
	fmt.Scanln(&id)

	fmt.Println("Enter Name For directory")
	scanner.Scan()
	d.name = scanner.Text()
	fmt.Println("Enter address For directory")
	scanner.Scan()
	d.address = scanner.Text()
	fmt.Println("Enter mobile For directory")
	fmt.Scanln(&d.mobile)
	fmt.Println("Enter pincode For directory")
	fmt.Scanln(&d.pincode)
	m[id] = d
}
func deletedetails() {
	var key int
	fmt.Println("Please Enter Vaild key to be deleted")
	fmt.Scan(&key)
	delete(m, key)

}

func getdetails() {
	var id int
	fmt.Println("Enter ID whose Details to be displayed")
	fmt.Scan(&id)
	fmt.Println(m[id])

}
func showdetails() {
	for key, value := range m {
		fmt.Println(key, value)
	}

}

func main() {
	for {
		var x int
		fmt.Println("Welcome in Telephone directory")
		fmt.Println("press 1 to find in directory")
		fmt.Println("press 2 to delete from directory")
		fmt.Println(" press 3  to list directory")
		fmt.Println("press 4 to add new entry")
		fmt.Println("Press -1 for Exit")
		fmt.Scan(&x)

		if x == 1 {
			getdetails()
		} else if x == 2 {
			deletedetails()
		} else if x == 3 {
			showdetails()
		} else if x == 4 {
			adddetails()
		} else {
			break
		}

	}
}
