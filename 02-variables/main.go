package main

import "fmt"

const PieValue = 70141414141441

func main() {
	var userName string = "Subrat Barik"
	var price float64= 253.5465634345345

	fmt.Println("Hello", userName)
	fmt.Printf("Type %T", userName)
	fmt.Printf("Float %T", price)

	// implicit type
	var website = "https://subratbarik/github.io/resume";
	fmt.Println(website)

	// no var style and only to be used inside a func
	noOfUsers := 30000 //walrus operator
	fmt.Println(noOfUsers)

	fmt.Println(PieValue)

}