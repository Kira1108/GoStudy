package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://www.cup.edu.cn")

	if err != nil {
		fmt.Println("Error request to baidu.")
		fmt.Println(err)
		os.Exit(1)
	}

	// Using read closer interface
	bs := make([]byte, 100000)
	n, _ := response.Body.Read(bs)
	fmt.Println("Read", n, "Bytes")
	fmt.Println(string(bs))

	// Using io.copy

	io.Copy(os.Stdout, response.Body)
}
