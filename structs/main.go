package main

import "fmt"

// Contract struct contains data of contact info
type ContactInfo struct {
	email   string
	address string
	zipcode int
}

// define a struct
// type Person struct {
// 	firstname string
// 	lastname  string
// 	ContactInfo   ContactInfo
// }

type Person struct {
	firstname string
	lastname  string
	// this equals to ContactInfo ContactInfo
	ContactInfo
}

// print infomation of a person
func (p Person) print() {
	fmt.Printf("<Person name = %v %v>\n",
		p.firstname,
		p.lastname)
	p.printContact()
}

// print contact info of a person
func (p Person) printContact() {
	fmt.Printf("ContactInfo: %v, %v, %v\n",
		p.ContactInfo.email,
		p.ContactInfo.address,
		p.ContactInfo.zipcode)
}

// update first name of a person
func (p *Person) updateFirstName(firstname string) {
	p.firstname = firstname
}

// update lastname of a person
func (p *Person) updateLastName(lastname string) {
	p.lastname = lastname
}

// update contact of a person
func (p *Person) updateContact(contactInfo ContactInfo) {
	p.ContactInfo = contactInfo
}

func main() {

	// Initialize a variable of type Person
	// and type Contract

	c := ContactInfo{
		email:   "327485253@163.com",
		address: "China Beijing",
		zipcode: 123442,
	}

	wanghuan := Person{
		firstname:   "wang",
		lastname:    "huan",
		ContactInfo: c,
	}
	wanghuan.print()

	wanghuan.updateFirstName("Laowang")
	wanghuan.updateLastName("badan")
	c1 := ContactInfo{
		email:   "ttfff@qq.com",
		address: "China Xinjiang",
		zipcode: 33445,
	}

	wanghuan.updateContact(c1)

	// declea a Person type and set attribute
	// var wanghuan Person
	// firstname and lastname got zero value -> "" and ""
	// wanghuan.print()

	// wanghuan.firstname = "wang"
	// wanghuan.lastname = "huan"

	wanghuan.print()
	// get and printing values
	// fmt.Println(wanghuan.firstname)
	// fmt.Println(wanghuan.lastname)
	// fmt.Println(wanghuan)
}
