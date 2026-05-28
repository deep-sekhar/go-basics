package main

import "fmt"
import "errors"
import "strings"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(intdiv(10, 2))
	// while printing use %s for string, %d for int, %f for float and %t for bool, %v for any type

	// return errors from functions and handle them in the caller function
	result, err := intdivWithError(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// switch cases
	switch {
	case result > 5:
		fmt.Println("Result is greater than 5")
	case result == 5:
		fmt.Println("Result is equal to 5")
	default:
		fmt.Println("Result is less than 5")
	}

	switch result {
	case 0:
		fmt.Println("Result is 0")
	case 1:
		fmt.Println("Result is 1")
	default:
		fmt.Println("Result is neither 0 nor 1")
	}

	// arrays
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)
	fmt.Println("1st 3 elements of array:", arr[0:3])

	// vectors in cpp is slices in go, which are dynamic arrays that can grow and shrink in size, and they are more commonly used than arrays in go
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
	slice = append(slice, 6)
	fmt.Println("Slice after appending 6:", slice)
	slice = append(slice, arr2[0:3]...) // append first 3 elements of arr2 to slice
	fmt.Println("Slice after appending first 3 elements of arr2:", slice)

	// maps in go are like unordered maps in cpp, which are collections of key-value pairs, where the keys are unique and the values can be of any type
	m := make(map[string]int)
	var m2 map[string]int = make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m2["four"] = 4
	fmt.Println("Map:", m)
	fmt.Println("Value for key 'two':", m["two"])
	delete(m, "two")
	fmt.Println("Map after deleting key 'two':", m)
	// check if key exists in map
	value, ok := m["two"]
	if ok {
		fmt.Println("Value for key 'two':", value)
	} else {
		fmt.Println("Key 'two' does not exist in map")
	}

	// iteration
	for name, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", name, value)
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, slice[i])
	}
	for i, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", i, value)
	}

	// we can create variables on the fly using := operator, while using things like for loops, if statements, switch cases, etc. and we can also use it to create variables in the main function
	if value, ok := m["one"]; ok {
		// executes if ok is true
		fmt.Println("Value for key 'one':", value)
	}

	// strings in go are immutable, which means that once a string is created, it cannot be changed, which is inefficient like in Java, but we can use strings.Builder to create a mutable string that can be modified without creating a new string every time, which is more efficient than using string concatenation with + operator
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(" ")
	sb.WriteString("World")
	fmt.Println("String Builder:", sb.String())

	// UTF-8 stores ASCII characters in 1 byte, but many non-ASCII characters take 2, 3, or 4 bytes.
	// So if you do s[i], you get the i-th byte, not the i-th character.
	// That is why non-ASCII text can look “broken” when you iterate a string byte-by-byte

	// But range loops over strings iterate over Unicode code points, so they work correctly with non-ASCII text.
	s := "Hi, 世界"
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %x\n", i, s[i])
	}
	for i, r := range s {
		fmt.Printf("Rune %d: %c\n", i, r)
	}

	// if you want random access to characters in a string, you can convert the string to a slice of runes, which are Go's type for Unicode code points, and then index into that slice. But be aware that this creates a new slice and copies the data, which can be inefficient for large strings.
	runes := []rune(s)
	fmt.Printf("Character at index 4: %c\n", runes[4])
}

func intdiv(a, b int) int {
	return a / b
}

// can return error as second return value, which is a built in type in go, and can be created using errors.New() function
func intdivWithError(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}