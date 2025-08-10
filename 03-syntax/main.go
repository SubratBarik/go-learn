package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Golang,"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	// comma ok || error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println(welcome, input)


}