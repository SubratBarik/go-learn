package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current Date & Time:")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-01-2006 01:02:05 Monday"))
}