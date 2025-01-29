package main

import (
    "fmt"
    "my-golang-app/pkg/utils"
)

func main() {
    fmt.Println("Welcome to My Golang App!")

    // Example usage of utility functions
    sum := utils.Add(5, 3)
    fmt.Printf("Sum: %d\n", sum)

    difference := utils.Subtract(5, 3)
    fmt.Printf("Difference: %d\n", difference)

    product := utils.Multiply(5, 3)
    fmt.Printf("Product: %d\n", product)
}