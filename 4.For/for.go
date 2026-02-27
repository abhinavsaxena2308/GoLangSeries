package main

import "fmt"

func main(){
	//while loop using for
	i:= 1
	for i<=3{
		fmt.Println(i)
		i=i+1
	}

	//infinite loop
// 	for {
// 		fmt.Println("1")
// 	}

	// for loop
	for i :=0;i<3;i++{
		println(i)
	}

	//range func
	for i:=range 3{
		println(i+1)
	}
}