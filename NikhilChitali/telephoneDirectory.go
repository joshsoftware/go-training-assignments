package main

import (
	"bufio"
	"fmt"
	"os"
)

type Person struct {
	id       int
	name     string
	address  string
	mobileNo string
	pinCode  string
}

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("############ Telephone Directory #############")

	dictionary := make(map[int]Person)

	recordId := 1

	for true {
		var option int
		fmt.Println("select operation:\npress 1 to find in directory\npress 2 to delete from directory\npress 3 to list directory\npress 4 to add new entry\npress 5 to exit\n")
		fmt.Scanln(&option)
		fmt.Println()
		if option == 5 {
			break
		}

		switch option {
		case 1:
			fmt.Println("Enter id whose details are required")
			var id int
			fmt.Scanln(&id)

			if record, doesExists := dictionary[id]; doesExists {
				fmt.Println(record)
			} else {
				fmt.Println("ID does not exist. Please Enter Valid ID")
			}

		case 2:
			fmt.Println("Enter ID of record that has to be deleted")
			var id int
			fmt.Scanln(&id)

			if _, doesExists := dictionary[id]; doesExists {
				delete(dictionary, id)
				fmt.Println("Record Deleted Successfully")
			} else {
				fmt.Println("Record does not exist")
			}

		case 3:
			if len(dictionary) == 0 {
				fmt.Println("No records in Dictionary")
			} else {
				fmt.Println("All Records that are in dictionary:")
				for record := range dictionary {
					fmt.Println(dictionary[record])
				}
			}
		case 4:

			var recordName string
			var recordAddress string
			var recordMobileNo string
			var recordPinCode string

			fmt.Println("Enter Name")
			scanner := bufio.NewScanner(os.Stdin)
			name := ""
			if scanner.Scan() {
				name += scanner.Text()
			}

			recordName = name

			fmt.Println("Enter Address")
			scanner = bufio.NewScanner(os.Stdin)
			name = ""
			if scanner.Scan() {
				name += scanner.Text()
			}

			recordAddress = name

			fmt.Println("Enter Mobile Numebr")
			fmt.Scan(&recordMobileNo)
			fmt.Println("Enter Pin Code")
			fmt.Scan(&recordPinCode)

			if len(recordMobileNo) != 10 || len(recordPinCode) != 6 {
				fmt.Println("Invalid Mobile Number or Pin Code!! Verify Again")
			} else {
				person := Person{recordId, recordName + ",", recordAddress + ",", recordMobileNo + ",", recordPinCode}

				dictionary[recordId] = person
				recordId++
				fmt.Println()
				fmt.Println("Record added Successfully!!!")
			}
		}
		fmt.Println()
		fmt.Println("############################")
	}

}
