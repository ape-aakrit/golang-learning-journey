package main

import "fmt"

func main() {
	// --------------------------
	// Conditional Statements in Go
	// --------------------------

	// 1. if statement
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult.")
	}

	// 2. if-else statement
	temperature := 30
	if temperature > 35 {
		fmt.Println("It's really hot today.")
	} else {
		fmt.Println("The temperature is comfortable.")
	}

	// 3. if-else if-else ladder
	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// 4. Short statement in if condition (variable scoped to if)
	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}

	// 5. switch statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}

	// 6. switch with multiple expressions in one case
	grade := "B"
	switch grade {
	case "A", "A+":
		fmt.Println("Excellent!")
	case "B", "B+":
		fmt.Println("Well done.")
	case "C":
		fmt.Println("You passed.")
	default:
		fmt.Println("Better luck next time.")
	}

	// 7. switch without an expression (like if-else)
	marks := 72
	switch {
	case marks >= 90:
		fmt.Println("Top scorer!")
	case marks >= 75:
		fmt.Println("Great job!")
	case marks >= 50:
		fmt.Println("You passed.")
	default:
		fmt.Println("Failed.")
	}
}
