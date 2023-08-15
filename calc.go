package main

import "fmt"

func add(x, y float64) float64 {
    return x + y
}

func subtract(x, y float64) float64 {
    return x - y
}

func multiply(x, y float64) float64 {
    return x * y
}

func divide(x, y float64) float64 {
    if y != 0 {
        return x / y
    }
    return 0 // Handle division by zero
}

func main() {
    var num1, num2 float64
    fmt.Print("Enter first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scan(&num2)

    fmt.Println("Addition:", add(num1, num2))
    fmt.Println("Subtraction:", subtract(num1, num2))
    fmt.Println("Multiplication:", multiply(num1, num2))
    fmt.Println("Division:", divide(num1, num2))
}
