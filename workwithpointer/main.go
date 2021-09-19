package main

// Value Types: int, float, string, bool, structs
// use pointer to change things in a function

// Reference Types: slice, map, channel, pointer, function
// don't worry about pointers with this

import "fmt"

type Person struct {
	firstname string
	lastname  string
}

func (p Person) print() {
	fmt.Printf("<Person name = %v %v>\n", p.firstname, p.lastname)
}

func (p Person) updateNoEffect(firstname string) {
	p.firstname = firstname
}

// *Person in function signature indicates a pointer type
// This is a receiver function on pointer to type Person
func (p *Person) updateNormal(firstname string) {
	// This * get the content of pointer p
	// *p gets a variable sitting in RAM, that has a type Person
	// and update its firstname to a new named specified by the function parameter.
	(*p).firstname = firstname
}

func (p *Person) updateSuperEasy(firstname string) {
	// This * get the content of pointer p
	// *p gets a variable sitting in RAM, that has a type Person
	// and update its firstname to a new named specified by the function parameter.
	p.firstname = firstname
}

func udpateSlice(s []string) {
	s[0] = "bye"
}

func main() {
	wanghuan := Person{firstname: "Wang", lastname: "Huan"}
	wanghuan.print()

	fmt.Println("\nUpdate but has no effect")
	wanghuan.updateNoEffect("Li")
	wanghuan.print()

	// get a pointer
	fmt.Println("\nSuccess update")
	// get the address of wanghuan
	// pointerToWanghuan is a pointer to a Person
	pointerToWanghuan := &wanghuan
	// use the receiver function of (pointer to Person)
	pointerToWanghuan.updateNormal("Li")
	wanghuan.print()

	// Tired of retrieve address from a variable
	fmt.Println("\nShortcut pointer receiver call")
	wanghuan.updateNormal("Niu")
	wanghuan.print()

	fmt.Println("\nSuper Stright forward pointer call")
	wanghuan.updateSuperEasy("Zhao")
	wanghuan.print()

	// 这就说明我们不需要取出一个东西的地址
	// 我们也不需要从一个pointertype中取出这个地址指向的东西
	// 一切简简单单，就能成功

	// But update slice works
	// 你把这个someSlice传给updateSlicefunction的时候，这个slice还是copy了
	// 但是copy过后的这个slice，和someSlice还是指向了同一个底层array
	// 所以这时候更改这个值就成功了
	someSlice := []string{"hi", "there", "how", "are", "you"}
	udpateSlice(someSlice)
	fmt.Println("\n", someSlice)
}
