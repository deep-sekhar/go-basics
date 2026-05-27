// go is statically typed language, which means that every variable has a type, and that type is known at compile time. This allows the compiler to catch type errors before the program is run, and it also allows the compiler to optimize the code for performance.

package main
import "fmt"

func main() {
	// data types in go
	var x int = 10
	fmt.Println(x)
	var y float64 = 3.14
	// same as double in c++, can be float32 or float64, but float64 is more common
	fmt.Println(y)
	var z string = "Hello, World!"
	fmt.Println(z)
	var b bool = true
	fmt.Println(b)

}