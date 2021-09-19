package main

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
}
