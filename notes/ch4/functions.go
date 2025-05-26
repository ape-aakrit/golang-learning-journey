package main

import "fmt"

// 1. Basic function with no parameters and no return value
func greet() {
	fmt.Println("Hello from a basic function!")
}

// 2. Function with parameters
func add(a int, b int) {
	fmt.Println("Sum is:", a+b)
}

// 3. Function with return value
func subtract(a int, b int) int {
	return a - b
}

// 4. Function with multiple return values
func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 5. Named return values (more readable)
func multiplyAndLabel(a int, b int) (product int, label string) {
	product = a * b
	label = "Result of multiplication"
	return // automatically returns named values
}

// 6. Variadic function (accepts multiple arguments)
func totalSum(nums ...int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

func main() {
	// Call basic function
	greet()

	// Call function with parameters
	add(10, 20)

	// Function with return value
	result := subtract(50, 15)
	fmt.Println("Subtraction result:", result)

	// Function with multiple return values
	q, r := divideAndRemainder(23, 5)
	fmt.Println("Quotient:", q, "Remainder:", r)

	// Named return values
	prod, msg := multiplyAndLabel(6, 7)
	fmt.Println(msg+":", prod)

	// Variadic function
	sum := totalSum(1, 2, 3, 4, 5)
	fmt.Println("Total sum:", sum)

	// 7. Anonymous function (used like a variable)
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Anonymous function result:", multiply(3, 4))

	// 8. Passing function as a variable
	doMath := func(f func(int, int) int, x int, y int) int {
		return f(x, y)
	}
	result2 := doMath(subtract, 20, 5)
	fmt.Println("Using function as parameter:", result2)
}
