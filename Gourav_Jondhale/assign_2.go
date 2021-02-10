package main
import (
	"bufio"
	"fmt"
	"os"
)

type Details struct{
	name,address string
	ph_no,pincode int
}

var m =make(map[int]Details)
var cnt int

func entry(d Details){
	
	fmt.Println("Enter name: ")
	var dummy string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		dummy=dummy+scanner.Text()
	}
	d.name=dummy+", "
	dummy=""
	// fmt.Scanln(&d.name)
	
	fmt.Println("Enter address: ")
	scanner = bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		dummy=dummy+scanner.Text()
	}
	d.address=dummy+", "
	dummy=""
	// fmt.Scanln(&d.address)

	fmt.Println("Enter phone number: ")
	fmt.Scanln(&d.ph_no)

	fmt.Println("Enter pincode: ")
	fmt.Scanln(&d.pincode)
	var digi,sec int
	sec=d.ph_no
	for{
		if(sec==0){
			break
		} else {
			digi++
			sec=sec/10
		}
	}
	if(digi!=10){
		fmt.Println()
		fmt.Println("=x=x=x=x=x=x=x=x=x=x=x=x=x=x=")
		fmt.Println("Phone number is invalid!!!!!")
		fmt.Println("=x=x=x=x=x=x=x=x=x=x=x=x=x=x=")
		fmt.Println()
	} else{
		cnt=cnt+1
	m[cnt]=d
	fmt.Println()
	fmt.Println("=================================")
	fmt.Println("Details recorded successfully!!!!")
	fmt.Println("=================================")
	fmt.Println()
	}
}
func delete_record(id int){

	var entity =m[id]
	if(entity.ph_no==0){
		fmt.Println()
		fmt.Println("========================")
		fmt.Println("Not a valid unique code")
		fmt.Println("========================")
		fmt.Println()
	} else {
		delete(m,id)
		cnt--
		fmt.Println()
		fmt.Println("===============================")
		fmt.Println("Record deleted successfully!!!!")
		fmt.Println("===============================")
		fmt.Println()
	}


}

func find_record(id int){
	var entity =m[id]
	if(entity.ph_no==0){
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("Please enter valid unique code")
		fmt.Println("==============================")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("========================")
		fmt.Println(m[id])
		fmt.Println("========================")
		fmt.Println()
	}

}

func main()  {
	var option int
	var d Details

	for {
		fmt.Println("Please enter the option to perform any option: ")
	fmt.Println("1----> to find record in directory")
	fmt.Println("2----> to delete any record from directory")
	fmt.Println("3----> to list directory")
	fmt.Println("4----> to add new entry in directory")
	fmt.Println("5----> to exit")
	fmt.Println("Your option goes here---->")
	fmt.Scanln(&option)
	fmt.Println()
	if option==5{
		break
	}	
	switch(option){
		case 1:{
			var id int
			fmt.Println("Please enter the unique code: ")
			fmt.Scanln(&id)
			find_record(id)
			fmt.Println()
		}
		case 2:{
			var id int
			fmt.Println("please enter the unique code: ")
			fmt.Scanln(&id)
			delete_record(id)
			fmt.Println()
			fmt.Println()
		}
		case 3:{
			fmt.Println()
			fmt.Println("====================")
			fmt.Println("All present records: ")
			// fmt.Println(m)
			for key,value :=range m{
				fmt.Println(key, value)
			}
			fmt.Println("=====================")
			fmt.Println()
		}
		case 4:{
			entry(d)
			fmt.Println()
		}
		case 5:
			break
	}
	
}

}