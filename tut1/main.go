// go is statically typed language, which means that every variable has a type, and that type is known at compile time. This allows the compiler to catch type errors before the program is run, and it also allows the compiler to optimize the code for performance.

package main
import "fmt"
import "unicode/utf8"

// initialise go module with go mod init <module name> and then run the code with go run main.go or build the code with go build main.go and then run the executable with ./main
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

	// declare variable either 
	// assign value
	var a = 5
	// mention type
	var c int
	// use := to declare and assign value without mentioning type and var keyword, but only inside functions
	d := "Go is awesome!"
	fmt.Println(a, c, d)

	// initialize multiple variables at once
	var i, j, k int = 1, 2, 3
	fmt.Println(i, j, k)

	// initialise string using " or ``
	s1 := "This is a string with double quotes"
	s2 := `This is a string with backticks, which can span multiple lines
	and can contain "quotes" without needing to escape them.`
	fmt.Println(s1)
	fmt.Println(s2)

	// len gives you number of bytes in the string i.e no of ASCII characters, not the number of characters in the string
	fmt.Println(len(s1))
	// incase non ASCII characters use utf8 package to get the number of characters in the string
	fmt.Println(utf8.RuneCountInString(s2))

	// default value of numbers is 0, for string is "" and for bool is false

	// package is collection of go files in the same directory that are compiled together, and can be imported by other go files to use the functions and variables defined in that package. The main package is the entry point of the program, and it must have a main function that is called when the program is run.

	// module is a collection of packages that are versioned together, and can be imported by other modules to use the packages defined in that module. A module is defined by a go.mod file, which specifies the module name and the dependencies of the module. The go.mod file is created by running go mod init <module name> in the terminal, and it is updated automatically when you add or remove dependencies from the module.
}