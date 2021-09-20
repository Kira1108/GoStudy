package main

import "fmt"

func main() {

	// create a nil map, assign will raise an error
	// var m map[string]string
	// m["Somekey"] = "somevalue"

	// correct way of creating an empty map
	m := make(map[string]string)
	m["somekey"] = "somevalue"
	m["otherkey"] = "othervalue"
	fmt.Println(m)

	// create an initialize map
	t := map[int]string{
		1: "Hello",
		2: "There",
	}
	fmt.Println(t)

	// assign key value pairs to a map
	m["somekey"] = "sengkou"
	fmt.Println(m)

	// delete a map key
	delete(m, "otherkey")
	fmt.Println(m)

	s := map[int]int{
		1: 2,
		2: 3,
		3: 5,
	}
	// loop through a map
	for k, v := range s {
		fmt.Println("key:", k, "value:", v)
	}

}
