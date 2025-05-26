package main

import (
	"fmt"
)

func main() {
	fmt.Println("🔁 Go Loops: All Variants You Need")

	////////////////////////////////
	// 1. Classic for loop (like C++)
	////////////////////////////////
	fmt.Println("\n1. Classic for loop")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}

	////////////////////////////////
	// 2. While-like loop
	////////////////////////////////
	fmt.Println("\n2. While-like loop (just 'for' with condition)")
	count := 1
	for count <= 3 {
		fmt.Println("count =", count)
		count++
	}

	////////////////////////////////
	// 3. Infinite loop (break manually)
	////////////////////////////////
	fmt.Println("\n3. Infinite loop with break")
	x := 1
	for {
		if x > 3 {
			break
		}
		fmt.Println("x =", x)
		x++
	}

	////////////////////////////////
	// 4. Using continue
	////////////////////////////////
	fmt.Println("\n4. Loop with continue (skip even numbers)")
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("odd i =", i)
	}

	////////////////////////////////
	// 5. Looping over slices using range
	////////////////////////////////
	fmt.Println("\n5. Range loop over slice")
	nums := []int{10, 20, 30, 40}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	////////////////////////////////
	// 6. Range over string (rune-wise)
	////////////////////////////////
	fmt.Println("\n6. Range over string (runes)")
	for i, char := range "Go💚Lang" {
		fmt.Printf("Index: %d, Rune: %q\n", i, char)
	}

	////////////////////////////////
	// 7. Range over map
	////////////////////////////////
	fmt.Println("\n7. Range over map")
	myMap := map[string]int{"Alice": 90, "Bob": 85}
	for key, value := range myMap {
		fmt.Printf("%s scored %d\n", key, value)
	}

	////////////////////////////////
	// 8. Nested loops
	////////////////////////////////
	fmt.Println("\n8. Nested loops")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}

	////////////////////////////////
	// 9. Using goto (not recommended unless needed)
	////////////////////////////////
	fmt.Println("\n9. Using goto")
	n := 0
JumpPoint:
	if n < 3 {
		fmt.Println("n =", n)
		n++
		goto JumpPoint
	}

	////////////////////////////////
	// 10. Loop Challenge: Sum of all odd numbers < 10
	////////////////////////////////
	fmt.Println("\n10. Loop Challenge: Sum of odd numbers < 10")
	sum := 0
	for i := 1; i < 10; i += 2 {
		sum += i
	}
	fmt.Println("Sum =", sum)
}
