package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No filename detected")
		os.Exit(1)
	}

	if strings.Split(os.Args[1], ".")[1] != "txt" {
		fmt.Println("Unable to read non-txt file")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Failed to read file")
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
