//Pragyesh Assignment

package main
import ("fmt")
//var m map[int] data
var ok bool


var id int=0
var v int=1
func main(){
	type data struct{
		name string
		email string
		pin int64
		number int64
	
	
	}
	  m:= make(map[int] data)
	 for ;v==1;{ 
	 fmt.Println("Press 1 to find")
	 fmt.Println("Press 2 to delete")
	 fmt.Println("Press 3 to list")
	 fmt.Println("Press 4 to add")
	 fmt.Println("Press 5 to exit")
	 var choice int
	 fmt.Scanln(&choice)
	 switch choice{
	   case 1:
		fmt.Println("Enter valid id")
		fmt.Scanln(&id)
		_,ok=m[id]
		if(ok){
			fmt.Print("ID:")
			fmt.Print(id)
			fmt.Print("  ")
			fmt.Print("NAME:")
			fmt.Print(m[id].name)
			fmt.Print("  ")
			fmt.Print("EMAIL:")
			fmt.Print(m[id].email)
			fmt.Print("  ")
			fmt.Print("PIN:")
			fmt.Print(m[id].pin)
			fmt.Print("  ")
			fmt.Print("NUMBER:")
			fmt.Println(m[id].number)
		 fmt.Println("Done")
		}
	     //else
		 // {fmt.Println("invlid id")  
	   case	2:
		fmt.Println("Enter valid id yoy want to delete")
		fmt.Scanln(&id)
		_,ok=m[id]
		if(ok){
		 delete(m, id)
		 fmt.Println("Deletion done")
		}
		//else{
		 //fmt.Println("invlid id")}
	   case 3:
		fmt.Println("Enter valid id you want to view detail")
		fmt.Scanln(&id)
		_,ok=m[id]
		if(ok){
			fmt.Print("ID:")
			fmt.Print(id)
			fmt.Print("  ")
			fmt.Print("NAME:")
			fmt.Print(m[id].name)
			fmt.Print("  ")
			fmt.Print("EMAIL:")
			fmt.Print(m[id].email)
			fmt.Print("  ")
			fmt.Print("PIN:")
			fmt.Print(m[id].pin)
			fmt.Print("  ")
			fmt.Print("NUMBER:")
			fmt.Println(m[id].number)}
		//else
		 // fmt.Println("invlid id")
	   case 4:
		fmt.Println("Enter name")
	   	var Name string
		var Mail string
		var Pin int64
		var Number int64 
		fmt.Scanln(&Name)
		fmt.Println("Enter mail")
		fmt.Scanln(&Mail)
		fmt.Println("Enter pin")
		fmt.Scanln(&Pin)
		fmt.Println("Enter number")
		fmt.Scanln(&Number)
		s:=data{name:Name,email:Mail,pin:Pin,number:Number}
		id=id+1
		m[id]=s
		
		fmt.Println("Adding Completed")
		fmt.Println("")

		
       case 5:
		v=0}

	   }
        




 }
