package main

import "fmt"

// a generic addition function that can add two values of any type that supports the + operator, and it returns the result of the addition
func add[T int | float64](a, b T) T {
	return a + b
}

func isEmpty[T any](s []T) bool {
	return len(s) == 0
}

// generics on struct
type gasEngine struct {
	fuelType string
}

type electricEngine struct {
	batteryCapacity int
}

type vehicle[T any] struct {
	name  string
	engine T
}

func main() {
	fmt.Println(add(5, 10))
	fmt.Println(add(5.5, 10.5))

	var a = make([]int, 0)
	fmt.Println(isEmpty(a)) // true

	vehicle1 := vehicle[gasEngine]{
		name: "Car 1",
		engine: gasEngine{
			fuelType: "Petrol",
		},
	}
	fmt.Println("Vehicle 1:", vehicle1)
}