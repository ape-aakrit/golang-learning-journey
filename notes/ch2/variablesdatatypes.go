package main

import "fmt"

func main() {
	// --------------------------
	// Variables and Data Types in Go
	// --------------------------

	// 1. Declaring variables with explicit type and initialization
	var age int = 25               // int = integer (whole numbers)
	var temperature float64 = 36.6 // float64 = floating point number (decimal numbers)
	var isStudent bool = true      // bool = boolean (true or false)
	var grade byte = 'A'           // byte = alias for uint8, used for ASCII chars

	// 2. Declaring variables with implicit type (type inferred by compiler)
	var name = "Alice" // string = sequence of characters (text)
	city := "New York" // short declaration (only inside functions)

	// 3. Multiple variable declaration
	var x, y, z int = 1, 2, 3 // multiple integers declared and initialized together

	// 4. Zero value variables (default values when not initialized)
	var uninitializedInt int       // defaults to 0
	var uninitializedFloat float64 // defaults to 0.0
	var uninitializedBool bool     // defaults to false
	var uninitializedString string // defaults to ""

	// 5. Constants (immutable values)
	const Pi = 3.14159 // constants cannot be changed after declaration

	// Printing variables and their types
	fmt.Println("Variables and their types:")
	fmt.Printf("age (int): %d\n", age)
	fmt.Printf("temperature (float64): %.2f\n", temperature)
	fmt.Printf("isStudent (bool): %t\n", isStudent)
	fmt.Printf("grade (byte): %c\n", grade)
	fmt.Printf("name (string): %s\n", name)
	fmt.Printf("city (string): %s\n", city)
	fmt.Printf("x, y, z (int): %d, %d, %d\n", x, y, z)

	fmt.Println("\nZero value variables:")
	fmt.Printf("uninitializedInt (int): %d\n", uninitializedInt)
	fmt.Printf("uninitializedFloat (float64): %.2f\n", uninitializedFloat)
	fmt.Printf("uninitializedBool (bool): %t\n", uninitializedBool)
	fmt.Printf("uninitializedString (string): '%s'\n", uninitializedString)

	fmt.Println("\nConstant:")
	fmt.Printf("Pi: %.5f\n", Pi)

	// 6. Using the var keyword without initialization
	var cityCode string
	cityCode = "NYC" // assign later
	fmt.Println("\nVariable declared without initialization:")
	fmt.Println("cityCode (string):", cityCode)
}
