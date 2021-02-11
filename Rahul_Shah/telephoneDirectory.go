package main

import (
	"fmt"
)

type entry struct{
	name string
	address string
	mobileNumber int
	pinCode int
}

func printInstructions(){
	fmt.Println("Enter 1 to find in directory")
	fmt.Println("Enter 2 to delete from directory")
	fmt.Println("Enter 3 to list directory")
	fmt.Println("Enter 4 to add new entry")
	fmt.Println("Enter 0 to exit")
}

func findEntry(dir map[int]entry){
	var id int
	fmt.Println(("Enter the id of the entry you want to find"))
	fmt.Scanln(&id)
	ent,exists := dir[id]
	if(exists){
		fmt.Println("Found the entry\n", ent)
	} else {
		fmt.Println("Couldn't find the entry, please try with another id")
		findEntry(dir)
	}
}

func addNewEntry()entry{
	var new entry
	fmt.Println("Enter name")
	fmt.Scanln(&new.name)
	fmt.Println("Enter address")
	fmt.Scanln(&new.address)

	getMobile:
	fmt.Println("Enter mobile number")
	fmt.Scanln(&new.mobileNumber)
	if(new.mobileNumber < 1000000000 || new.mobileNumber > 9999999999){
		fmt.Println("Please enter a valid 10 digit mobile number")
		goto getMobile
	}

	getPinCode:
	fmt.Println("Enter pin code")
	fmt.Scanln(&new.pinCode)
	if(new.pinCode < 100000 || new.pinCode > 999999){
		fmt.Println("Please enter a valid 6 digit pin code")
		goto getPinCode
	}

	return new
}

func deleteEntry(dir map[int]entry){
	var id int
	fmt.Println("Enter the id of the entry to delete:")
	fmt.Scanln(&id)
	_,exists:=dir[id]
	if(exists){
		delete(dir, id)
	} else {
		fmt.Println("No such entry found")
	}
}

func main(){
	id:=0
	directory := make(map[int]entry)
	directory[id] = entry{"Rahul", "Pune", 9820122589, 400019}
	id++
	fmt.Println(directory)
	for{
		printInstructions()
		var input int
		fmt.Scanln(&input)
		switch(input){
			case 1:
				findEntry(directory)
			case 2:
				deleteEntry(directory)
			case 3:
				for k,v:=range directory{
					fmt.Printf("%v: %v \n", k,v)
				}
			case 4:
				directory[id] = addNewEntry()
			case 0:
				return
		}	

	}
}