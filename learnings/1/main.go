package main

import (
	"fmt" //pkg imported from the std lib of go
	// "time"
)

var name = "Khushi";
const PI = 3.14;
//constants grouping for multiple configs 
const (
	PORT = 8080;
	host = "localhost"
)
func main() {
	//shorthand syntax => should always be written in the scope of a function
	// age := 22;
	// fmt.Println("Hello world");
	// fmt.Println(name);
	// fmt.Println(age);
	// fmt.Print(PORT,host);

	//while loop in go using for
	// i := 0;
	// for i <= 4{
	// 	fmt.Println(i);
	// 	i+=1;
	// }

	//for loop
	// for i:=0;i<=4;i++{
	// 	if(i == 1){
	// 		continue;
	// 	}
	// 	fmt.Print(i);
	// }

	//multiple cases switch case
	// switch time.Now().Weekday(){
	// 	case time.Saturday, time.Sunday:
	// 		fmt.Println("weekend");
	// 	default:
	// 		fmt.Println("weekday");
	// }

	//type switch
	//interface{} => an empty interface indicates any type
	// myType := func(i any){
	// 	switch i.(type){
	// 		case int:
	// 			fmt.Println("integer");
	// 		case string:
	// 			fmt.Println("string");
	// 	}
	// }

	// myType(22);

	//arrays
	
	var nums [4]int;
	nums[0] = 1;
	fmt.Println(nums);

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b);
	
	//go can infer the size of the array using ...

	c := [...]int{9,8,6,4};
	fmt.Println(c);
	fmt.Println(len(c));

}

