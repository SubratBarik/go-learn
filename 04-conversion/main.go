package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    fmt.Println("Barik Da's Pizza")
    fmt.Print("Please rate our service (1-5): ")

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    trimmed := strings.TrimSpace(input)

    numRating, err := strconv.ParseFloat(trimmed, 64)
    if err != nil {
        fmt.Println("Invalid input. Please enter a number.")
        return
    }

    fmt.Println("Thanks for rating us:", numRating)
    fmt.Println("Added 1 to the rating:", numRating+1)
}
