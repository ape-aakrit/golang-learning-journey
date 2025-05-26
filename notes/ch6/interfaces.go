package main

import (
	"fmt"
	"math"
)

///////////////////////////
// 1. Define Interface
///////////////////////////

type Shape interface {
	Area() float64
	Perimeter() float64
}

///////////////////////////
// 2. Shape Structs
///////////////////////////

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

///////////////////////////
// 3. Implementing Interface
///////////////////////////

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return 4 * s.side
}

///////////////////////////
// 4. Parameterized Function
///////////////////////////

func printShapeDetails(shape Shape) {
	fmt.Printf("Shape: %T\n", shape)
	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
	fmt.Println("--------------------------")
}

///////////////////////////
// 5. Main Program with User Input
///////////////////////////

func main() {
	var choice int

	fmt.Println("📐 Shape Calculator using Interfaces")
	fmt.Println("1. Rectangle")
	fmt.Println("2. Circle")
	fmt.Println("3. Square")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		var w, h float64
		fmt.Print("Enter width: ")
		fmt.Scanln(&w)
		fmt.Print("Enter height: ")
		fmt.Scanln(&h)
		rect := Rectangle{width: w, height: h}
		printShapeDetails(rect)

	case 2:
		var r float64
		fmt.Print("Enter radius: ")
		fmt.Scanln(&r)
		circ := Circle{radius: r}
		printShapeDetails(circ)

	case 3:
		var s float64
		fmt.Print("Enter side: ")
		fmt.Scanln(&s)
		sq := Square{side: s}
		printShapeDetails(sq)

	default:
		fmt.Println("Invalid choice. Please run the program again.")
	}
}
