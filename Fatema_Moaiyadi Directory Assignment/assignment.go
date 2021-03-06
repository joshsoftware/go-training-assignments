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
	pincode int32
}

var m = make(map[int]details)

var id = 0

func addNew(id int) {
	var d details
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter name:")
	scanner.Scan()
	d.name = scanner.Text()
	fmt.Println("Enter address:")
	scanner.Scan()
	d.address = scanner.Text()
	fmt.Println("Enter mobile number:")
	fmt.Scanln(&d.mobile)
	fmt.Println("Enter pincode:")
	fmt.Scanln(&d.pincode)
	m[id] = d
}

func getDetails(id1 int) {
	if m[id1] != 0 {
		fmt.Println("Details corresponding to this id are:")
		fmt.Println("Name=", m[id1].name)
		fmt.Println("Address=", m[id1].address)
		fmt.Println("Mobile=", m[id1].mobile)
		fmt.Println("Pincode=", m[id1].pincode)
	} else {
		fmt.Println("Id does not exist!")
	}

}

func del(id1 int) {
	delete(m, id1)
}

func show() {
	for key, value := range m {
		fmt.Println("id=", key)
		fmt.Println("Name=", value.name)
		fmt.Println("Address=", value.address)
		fmt.Println("Mobile=", value.mobile)
		fmt.Println("Pincode=", value.pincode)
	}
}
func main() {
	var option int
	for {
		fmt.Println("1. Get details with id")
		fmt.Println("2. Add new entry")
		fmt.Println("3. Delete from directory")
		fmt.Println("4. List from the directory")

		fmt.Print("Enter your choice:")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var id1 int
			fmt.Print("Enter id to get details:")
			fmt.Scanln(&id1)
			getDetails(id1)
		case 2:
			id++
			addNew(id)
		case 3:
			var id1 int
			fmt.Print("Enter id to delete:")
			fmt.Scanln(&id1)
			del(id1)
		case 4:
			show()
		}
		var ans string
		fmt.Print("Do you want to continue? (Y/N)")
		fmt.Scanln(&ans)
		if ans == "N" {
			break
		} else {
			continue
		}
	}
}
