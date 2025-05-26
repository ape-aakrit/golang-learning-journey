package main

import "fmt"

// 1. Defining a basic struct
type Person struct {
	name string
	age  int
	city string
}

// 2. Nested struct
type ContactInfo struct {
	email string
	phone string
}

type Employee struct {
	id      int
	details Person      // Nested struct
	contact ContactInfo // Another nested struct
}

// 3. Method on a struct (receiver function)
func (p Person) greet() {
	fmt.Printf("Hi, my name is %s and I live in %s.\n", p.name, p.city)
}

// 4. Method with pointer receiver (to modify original struct)
func (p *Person) updateCity(newCity string) {
	p.city = newCity
}

func main() {
	// Creating a struct instance using field names
	p1 := Person{name: "Alice", age: 25, city: "New York"}
	fmt.Println("Person 1:", p1)

	// Calling a method on struct
	p1.greet()

	// Using pointer receiver method to modify city
	p1.updateCity("San Francisco")
	fmt.Println("Updated Person 1:", p1)

	// Creating struct without field names (not recommended)
	p2 := Person{"Bob", 30, "Chicago"}
	fmt.Println("Person 2:", p2)

	// Struct pointer
	p3 := &Person{name: "Charlie", age: 28, city: "Seattle"}
	p3.updateCity("Austin")
	fmt.Println("Pointer struct updated:", *p3)

	// Nested struct
	emp := Employee{
		id:      101,
		details: Person{name: "Diana", age: 32, city: "Boston"},
		contact: ContactInfo{email: "diana@example.com", phone: "1234567890"},
	}
	fmt.Println("Employee:", emp)
	fmt.Println("Employee Name:", emp.details.name)
	fmt.Println("Employee Email:", emp.contact.email)

	// Anonymous struct (used when you don't need a named type)
	student := struct {
		name  string
		grade string
	}{
		name:  "Eva",
		grade: "A+",
	}
	fmt.Println("Anonymous struct:", student)
}
