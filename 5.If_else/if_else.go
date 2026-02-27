package main

import "fmt"

func main() {
	i := 12
	if i < 10 {
		println("number is less than 10")
	} else {
		println("number is greater than 10")
	}

	//logical opn
	var role = "user"
	var hasPermissions = false

	if role == "admin" || hasPermissions {
		fmt.Println("admin verified")
	} else {
		println("admin not verified")
	}

	//shorter method
	// we can declare a variable inside if construct
	if age:=15; age>=18{
		fmt.Println("person is adult")
	}else{
		fmt.Println("person is an adult")
	}	

	// go does not have ternary , you will have to use normal if else
}