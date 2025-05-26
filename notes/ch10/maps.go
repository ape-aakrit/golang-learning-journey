/*
📘 Go Maps — Full Notes (with Comments)
🔑 What is a Map?
A map is an unordered collection of key-value pairs.

It’s like a dictionary in Python or unordered_map in C++.

Syntax: map[keyType]valueType

✅ Declaring and Initializing Maps
a. Using make():
go
Copy code
m := make(map[string]int) // map from string to int
b. Using map literals:
go
Copy code

	m := map[string]string{
	    "name": "Alice",
	    "city": "Delhi",
	}

✅ Adding / Updating Elements
go
Copy code
m["age"] = 25      // Add or update key "age"
✅ Accessing Elements
go
Copy code
val := m["name"]   // Gets the value
✅ Checking if a Key Exists
go
Copy code
val, ok := m["name"]

	if ok {
	    fmt.Println("Key exists:", val)
	} else {

	    fmt.Println("Key does not exist")
	}

✅ Deleting a Key
go
Copy code
delete(m, "name")
✅ Length of a Map
go
Copy code
len(m) // returns number of key-value pairs
✅ Iterating Over a Map
go
Copy code

	for key, value := range m {
	    fmt.Println(key, value)
	}

✅ Maps are Reference Types
Maps are reference types → passed by reference.

Changing a map in one function affects it everywhere.

⚠️ Zero Value of Map
go
Copy code
var m map[string]int
// m is nil → panic if accessed without initialization
Use make() before using it.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("🗺️ Exploring Maps in Go")

	// 1. Creating a map using make()
	studentAge := make(map[string]int)
	studentAge["Alice"] = 21
	studentAge["Bob"] = 23
	fmt.Println("1. Map using make():", studentAge)

	// 2. Creating a map using a literal
	countryCapital := map[string]string{
		"India":   "New Delhi",
		"USA":     "Washington D.C.",
		"Germany": "Berlin",
	}
	fmt.Println("2. Map literal:", countryCapital)

	// 3. Accessing values
	fmt.Println("3. Capital of India:", countryCapital["India"])

	// 4. Checking if a key exists
	val, exists := countryCapital["USA"]
	if exists {
		fmt.Println("4. Key exists! Capital of USA:", val)
	} else {
		fmt.Println("4. Key does not exist")
	}

	// 5. Accessing a non-existing key
	val2, ok := countryCapital["Mars"]
	fmt.Println("5. Accessing non-existing key → Value:", val2, "Exists?", ok)

	// 6. Deleting a key
	delete(countryCapital, "Germany")
	fmt.Println("6. After deleting Germany:", countryCapital)

	// 7. Map length
	fmt.Println("7. Number of entries:", len(countryCapital))

	// 8. Iterating over map
	fmt.Println("8. Iterating countryCapital map:")
	for country, capital := range countryCapital {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}

	// 9. Map is reference type
	newMap := studentAge
	newMap["Charlie"] = 25
	fmt.Println("9. Original studentAge map after modifying newMap:", studentAge)

	// 10. Nil map (not initialized)
	var uninitMap map[string]int
	fmt.Println("10. Nil map value:", uninitMap)
	// uninitMap["test"] = 1 // ❌ This will panic: assignment to entry in nil map

	// 11. Nested maps
	userData := map[string]map[string]string{
		"user1": {
			"name":  "Alice",
			"email": "alice@example.com",
		},
		"user2": {
			"name":  "Bob",
			"email": "bob@example.com",
		},
	}
	fmt.Println("11. Nested map (userData):", userData)
	fmt.Println("   user1's email:", userData["user1"]["email"])
}

/*
OUTPUT
🗺️ Exploring Maps in Go
1. Map using make(): map[Alice:21 Bob:23]
2. Map literal: map[Germany:Berlin India:New Delhi USA:Washington D.C.]
3. Capital of India: New Delhi
4. Key exists! Capital of USA: Washington D.C.
5. Accessing non-existing key → Value:  Exists? false
6. After deleting Germany: map[India:New Delhi USA:Washington D.C.]
7. Number of entries: 2
8. Iterating countryCapital map:
Country: USA, Capital: Washington D.C.
Country: India, Capital: New Delhi
9. Original studentAge map after modifying newMap: map[Alice:21 Bob:23 Charlie:25]
10. Nil map value: map[]
11. Nested map (userData): map[user1:map[email:alice@example.com name:Alice] user2:map[email:bob@example.com name:Bob]]
   user1's email: alice@example.com
*/
