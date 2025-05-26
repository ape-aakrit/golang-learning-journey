package main

import (
	"errors"
	"fmt"
)

/////////////////////////
// 1. Division with Error Handling
/////////////////////////

// divide takes two numbers and returns result or an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("❌ cannot divide by zero")
	}
	return a / b, nil
}

/////////////////////////
// 2. Advanced: Square Root Function with Custom Error
/////////////////////////

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("❌ cannot take square root of negative number: %v", n)
	}
	return n * n, nil
}

/////////////////////////
// 3. Main with Input + Panic Example
/////////////////////////

func main() {
	var x, y float64
	fmt.Print("Enter numerator: ")
	fmt.Scanln(&x)

	fmt.Print("Enter denominator: ")
	fmt.Scanln(&y)

	result, err := divide(x, y)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("✅ Result of %.2f / %.2f = %.2f\n", x, y, result)
	}

	/////////////////////////
	// Square Root Example
	/////////////////////////

	var num float64
	fmt.Print("\nEnter number to square: ")
	fmt.Scanln(&num)

	sq, err := squareRoot(num)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("✅ Square of %.2f is %.2f\n", num, sq)
	}

	/////////////////////////
	// Panic Example (used rarely)
	/////////////////////////

	fmt.Println("\n⚠️ Demonstrating panic...")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("⚡ Recovered from panic:", r)
		}
	}()

	triggerPanic() // This will panic
}

func triggerPanic() {
	panic("🔥 something went very wrong!")
}
