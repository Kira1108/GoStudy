package main

import "fmt"

// define a struct
type Person struct {
	firstname string
	lastname  string
}

func (p Person) print() {
	fmt.Printf("<Person name = %v %v>\n", p.firstname, p.lastname)
}

func main() {

	// Initialize a variable of type Person
	// wanghuan := Person{
	// 	firstname: "wang",
	// 	lastname:  "huan",
	// }

	// declea a Person type and set attribute
	var wanghuan Person
	// firstname and lastname got zero value -> "" and ""
	wanghuan.print()

	wanghuan.firstname = "wang"
	wanghuan.lastname = "huan"

	wanghuan.print()
	// get and printing values
	// fmt.Println(wanghuan.firstname)
	// fmt.Println(wanghuan.lastname)
	// fmt.Println(wanghuan)
}
