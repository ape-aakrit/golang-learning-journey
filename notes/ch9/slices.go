/*
🧠 Go Slices — Full Notes (With Comments)
A slice is a flexible, powerful, and lightweight abstraction over an array.

✅ 1. What is a Slice?
A slice is a dynamically-sized view into the elements of an array.

Unlike arrays, slices don’t have a fixed size.

Syntax: []T where T is the type.

go
Copy code
var s []int // nil slice (has length 0, capacity 0)
✅ 2. Creating Slices
a. From an array:
go
Copy code
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Elements 2, 3, 4
b. Using make():
go
Copy code
s := make([]int, 5)        // length 5, capacity 5
s2 := make([]int, 5, 10)   // length 5, capacity 10
c. Using a slice literal:
go
Copy code
s := []int{10, 20, 30}
✅ 3. Slice Properties
go
Copy code
len(s)     // number of elements
cap(s)     // underlying array capacity from starting index
✅ 4. Slice Internals
A slice is a struct:

go
Copy code

	type slice struct {
	    ptr *array
	    len int
	    cap int
	}

✅ 5. Modifying Slices
Slices are references to underlying arrays.

Modifying a slice modifies the array.

go
Copy code
s := []int{1, 2, 3}
s[1] = 99 // Now s = [1 99 3]
✅ 6. Slicing a Slice
go
Copy code
s := []int{10, 20, 30, 40, 50}
s2 := s[1:4] // [20, 30, 40]
You can do:

s[:] → full slice

s[:n] → up to n

s[n:] → from n onward

✅ 7. Appending to a Slice
go
Copy code
s := []int{1, 2, 3}
s = append(s, 4, 5) // [1, 2, 3, 4, 5]
When capacity is exceeded, Go allocates a new underlying array.

✅ 8. Copying Slices
go
Copy code
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
✅ 9. Multi-dimensional Slices
go
Copy code

	matrix := [][]int{
	    {1, 2},
	    {3, 4},
	}

✅ 10. Nil vs. Empty Slice
Nil: var s []int → s == nil is true

Empty: s := []int{} → not nil, but len(s) == 0
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("🔪 Go Slices In-Depth")

	// 1. Creating a slice from an array
	arr := [5]int{10, 20, 30, 40, 50}
	s1 := arr[1:4] // [20, 30, 40]
	fmt.Println("1. Slice from array:", s1)

	// 2. Slice literals
	s2 := []string{"Go", "Rust", "Python"}
	fmt.Println("2. Slice literal:", s2)

	// 3. Using make()
	s3 := make([]int, 3)    // [0 0 0], len=3, cap=3
	s4 := make([]int, 3, 5) // [0 0 0], len=3, cap=5
	fmt.Println("3. make() slices:", s3, s4)

	// 4. Length and Capacity
	fmt.Println("4. len:", len(s4), "cap:", cap(s4))

	// 5. Slicing a slice
	s5 := []int{1, 2, 3, 4, 5}
	subslice := s5[1:4] // [2 3 4]
	fmt.Println("5. Sub-slice:", subslice)

	// 6. Appending
	s6 := []int{1, 2}
	s6 = append(s6, 3, 4)
	fmt.Println("6. Appended slice:", s6)

	// 7. Appending another slice
	s7 := []int{5, 6}
	s6 = append(s6, s7...) // spread operator
	fmt.Println("7. Appended another slice:", s6)

	// 8. Modifying a slice modifies underlying array
	fmt.Println("8. Before modifying s1:", s1)
	s1[0] = 999
	fmt.Println("   After modifying s1:", s1)
	fmt.Println("   Underlying array arr:", arr)

	// 9. Copy slices
	original := []int{100, 200, 300}
	copied := make([]int, len(original))
	copy(copied, original)
	fmt.Println("9. Copied slice:", copied)

	// 10. Nil vs Empty
	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("10. Nil slice? ", nilSlice == nil)
	fmt.Println("    Empty slice? len:", len(emptySlice), "cap:", cap(emptySlice))

	// 11. Multidimensional slice
	matrix := [][]int{
		{1, 2},
		{3, 4, 5},
		{},
	}
	fmt.Println("11. Matrix (2D slice):", matrix)

	// 12. Loop over slice using range
	fmt.Println("12. Looping over slice:")
	for i, v := range s6 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
