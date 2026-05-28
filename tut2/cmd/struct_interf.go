// learn about struct and interface
package main

import "fmt"

type Person struct {
	name string
	age  int
}

// define a method for the Person struct, which is a function that has a receiver of type Person, and it can be called on any variable of type Person
func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

type Greeter interface {
	greet()
}

// concrete 

func main() {
	p := Person{name: "Alice", age: 30}
	p.greet()

	// can define and initlise a struct and call its method in one line
	var p2 = Person{name: "Bob", age: 25}
	p2.greet()

	// interfaces in go are a way to define a set of methods that a type must implement, and they are used to achieve polymorphism in go
	// no explicit "implements" keyword is needed in go, if a type has all the methods defined in an interface, then it automatically implements that interface
	var g Greeter = p
	g.greet()
}