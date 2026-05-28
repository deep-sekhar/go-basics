// learn about pointers in go

package main

import "fmt"

func main() {
	var p *int // declare a pointer to an int
	fmt.Println("Pointer p:", p) // prints <nil> because p is not initialized
	p = new(int)
	// print value and address of pointer p
	fmt.Println("Address of p:", &p) // prints the address of pointer p
	fmt.Println("Value of p:", p) // prints the address of the int that p points to
	fmt.Println("Value at p:", *p) // prints 0 because the int that p points to is initialized to 0
	
	// if we dont assign a value to the int that p points to, it will be initialized to 0 by default, and we can assign a value to it using the dereference operator *
	*p = 42
	fmt.Println("Value at p after assignment:", *p) // prints 42

	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	arr2 := arr // creates a new array that is a copy of arr
	slice2 := slice // creates a new slice that is a reference to the same underlying array as slice
	arr2[0] = 10
	slice2[0] = 10
	fmt.Println("Original array:", arr) // prints [1 2 3 4 5] because arr is a value type and arr2 is a copy of arr
	fmt.Println("Original slice:", slice) // prints [10 2 3 4 5] because slice is a reference type and slice2 is a reference to the same underlying array as slice
}

// in functions instead of arrays we can pass a pointer to an array, and instead of slices we can pass a pointer to a slice, and we can modify the original array or slice in the caller function using the pointer
func modifyArray(arr *[5]int) {
	(*arr)[0] = 10 // modify the first element of the array that arr points to
}

func modifySlice(slice *[]int) {
	(*slice)[0] = 10 // modify the first element of the slice that slice points to
}

// Clearing pass by ref and pass by value confusion:
// in Java we only have pass by value
// objects like Map, ArrayList pass by value but the value is a reference to the object, so when we modify the object in the caller function, it modifies the original object. In general also when we create objects they dont hold values but references to the values.
// But some objects like String, Integer, Float pass by value and they are immutable, so when we modify them in the caller function, it creates a new object and does not modify the original object.

// BUT cpp has true pass by ref and pass by value too.
// IN CPP also when we pass a pointer to a function, we are passing a copy of the pointer, but the pointer itself is a reference to the original variable, so when we modify the variable that the pointer points to in the caller function, it modifies the original variable. 
// But since its just a ref to original variable say if I set it to ANOTHER pointer like NULL, then it will not modify the original variable.
// In cpp true pass by reference is achieved using reference variables, which are aliases for the original variable. We need to use & keyword for that. Eg:
// void modify(int &x) {} or void modify(int *&x) {} for pointer reference
// Note : & in front of a variable is used to get the address of the variable, and & in front of a type in a function parameter is used to declare a reference variable.

// In Go we only have pass by value, but the value can be a reference to an object, so when we modify the object in the caller function, it modifies the original object just like in Java.
// Things which are passed by copy which are ref to objects are s
// slices, maps, channels, interfaces, functions and pointers
// So when we pass them to a function, we are passing a copy of the reference to the object, but since the reference points to the same underlying object, when we modify the object in the caller function, it modifies the original object.
// Note: Although we have & in Go , its only used to get the address of a variable, and we dont have reference variables like in cpp, so we cannot achieve true pass by reference in Go. We can only achieve pass by value with reference types.
