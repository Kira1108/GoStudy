package main

import "fmt"

// define a function set to conform to
type Bot interface {
	getGreeting() string
}

type SomeType interface {
	somefunc1(string, int) (string error) // take in string and int, return string and error
	somefunc2()                           // take in nothing, return nothing
	somefunc3() string                    // take in nothing return string
	// somefunc(type1, type2) no actual variable name inside this function declaration
}

// a type englishBot
type englishBot struct {
}

// a type spanishBot
type spanishBot struct {
}

// receiver on an interface type is not valid
// func (b Bot) printGreetingR() {
// 	fmt.Println(b.getGreeting())
// }

// english bot greeting content
func (englishBot) getGreeting() string {
	// you did't use the englishBot variable
	// so ignore the variable name in receiver statement (eb englishBot)
	return "Hello"
}

// spanish bot greeting content
func (spanishBot) getGreeting() string {
	return "Hola"
}

// english bot print greeting
// func (e englishBot) printGreeting() {
// 	fmt.Println(e.getGreeting())
// }

// spanish bot print greeting
// func (s spanishBot) printGreeting() {
// 	fmt.Println(s.getGreeting())
// }

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	e := englishBot{}
	printGreeting(e)

	s := spanishBot{}
	printGreeting(s)
}
