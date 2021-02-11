// Build an interactive a telephone directory with id(has to be unique), name, address,
// mobile number, pin code.
// 1.Should be able to get details with id
// 2.Add new to directory
// 3.Delete from directory
// 4.List directory

package main

import "fmt"

type contact struct {
	name          string
	address       string
	mobile_number string
	pincode       string
}

func find_in_directory(directory map[int]contact, id int) {
	fmt.Println("Name:", directory[id].name, "\nAddress:", directory[id].address, "\nMobile Number:", directory[id].mobile_number, "\nPincode:", directory[id].pincode)
}

func delete_from_directory(directory map[int]contact, id int) {
	delete(directory, id)
	fmt.Println(directory)
}

func list_directory(directory map[int]contact) {
	fmt.Println(directory)
}

func add_to_directory(name, address, mobile_number, pincode string, directory map[int]contact) contact {
	// add value to map
	p := contact{
		name:          name,
		address:       address,
		mobile_number: mobile_number,
		pincode:       pincode,
	}
	return p
}

func main() {
	directory := make(map[int]contact)
	end_loop := 0
	id := 0
	var name, address, mobile_number, pincode string
	var input_id int
	var option int
	for end_loop == 0 {
		fmt.Println("Press \n1 to find in directory \n2 to delete from directory \n3 to list directory \n4 to add new entry \n5 to exit")
		fmt.Scan(&option)
		switch option {
		case 1:
			// fmt.Println("one, find")
			fmt.Println("Enter the id to find")
			fmt.Scan(&input_id)
			find_in_directory(directory, input_id)
		case 2:
			// fmt.Println("two, delete")
			fmt.Println("Enter the id to delete")
			fmt.Scan(&input_id)
			delete_from_directory(directory, input_id)
		case 3:
			// fmt.Println("three, List")
			list_directory(directory)
		case 4:
			// fmt.Println("four, Add")
			fmt.Println("Enter space seperated name, address, mobile_number, pincode")
			fmt.Scan(&name, &address, &mobile_number, &pincode)
			directory[id] = add_to_directory(name, address, mobile_number, pincode, directory)
			id += 1
		case 5:
			fmt.Println("Exiting")
			end_loop = 1
		default:
			fmt.Println("Please enter a valid option")
		}
	}
}

/*
akansha@akansha-Latitude-3490:~/golang/assignments$ go run AkanshaKumari/telephone_directory.go
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
3
map[]
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
4
Enter space seperated name, address, mobile_number, pincode
Akansha Pune 989887676 410506
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
3
map[0:{Akansha Pune 989887676 410506}]
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
1
Enter the id to find
0
Name: Akansha
Address: Pune
Mobile Number: 989887676
Pincode: 410506
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
2
Enter the id to delete
0
map[]
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
3
map[]
Press
1 to find in directory
2 to delete from directory
3 to list directory
4 to add new entry
5 to exit
5
Exiting
akansha@akansha-Latitude-3490:~/golang/assignments$
*/
