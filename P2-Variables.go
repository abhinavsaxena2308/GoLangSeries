package main //it defines the main package
import "fmt" //it imports the fmt package for formatted I/O

func main() {
	var name string = "abhinav"
	name1 := "hello"
	var age int = 21
	age1 := 21
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}
	var isStudent bool = true
	fmt.Println(name, name1, age, age1, isStudent)

}
