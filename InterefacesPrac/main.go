package main

import "fmt"

// declaring an interface shape
type shape interface{
	getArea() float64
}

// declaring custom struct types
type square struct{
	sideLength float64
}

type triangle struct{
	height float64
	base float64
}

type y struct{}

func main(){
	// Initializing the structs
	sq := square{
		sideLength: 10.0,
	}

	tri := triangle{
		height: 12.0,
		base: 7.0,
	}

	// print the area of the shapes
	printArea(sq)
	printArea(tri)

}

// function to print the area
func printArea(s shape){
	fmt.Println(s.getArea())
}

// getting the area of the square
func (s square) getArea() float64{
	return s.sideLength * s.sideLength
}

// getting the area of the triangle
func (t triangle) getArea() float64{
	return 0.5 * t.base * t.height
}
