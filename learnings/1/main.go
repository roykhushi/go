package main

import "fmt" //pkg imported from the std lib of go

var name = "Khushi";
const PI = 3.14;
//constants grouping for multiple configs 
const (
	PORT = 8080;
	host = "localhost"
)
func main() {
	//shorthand syntax => should always be written in the scope of a function
	age := 22;
	fmt.Println("Hello world");
	fmt.Println(name);
	fmt.Println(age);
	fmt.Print(PORT,host);
}