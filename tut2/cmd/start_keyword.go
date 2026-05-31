package main

// the make keyword is used to create slices, maps and channels in go, and it is used to initialize the internal data structures of these types, and it returns a value of the specified type
// for slices, make creates a slice with a specified length and capacity, and it initializes the underlying array with zero values, and it returns a slice that points to the underlying array
// for maps, make creates a map with a specified initial capacity, and it initializes the internal data structures of the map, and it returns a map that is ready to use
// for channels, make creates a channel with a specified buffer size, and it initializes the internal data structures of the channel, and it returns a channel that is ready to use

func main() {
	// create a slice of ints with length 5 and capacity 10 using make
	slice := make([]int, 5, 10)
	fmt.Println("Slice:", slice)
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))
	slice = append(slice, 1, 2, 3, 4, 5) // here we are appending 5 more elements to the slice, which will increase its length to 10 and it will not need to allocate a new underlying array because the capacity is already 10
	// similar to creating a vector of ints with initial size 5 and capacity 10 in cpp using vector<int> v(5); v.reserve(10);

	// create a map of string to int with initial capacity 10 using make
	m := make(map[string]int, 10)
	m["one"] = 1
	m["two"] = 2
	fmt.Println("Map:", m)
	// this is intersting because creating a map without make will result in a nil map, and trying to assign a value to a key in a nil map will result in a runtime panic. In simple words without make, the map points to nil/null.
	// But in cpp when we create an unordered_map without specifying the initial capacity, it will still be created in the stack and it will point in the heap to the internal data structures of the map. When we insert new data, it will allocate memory in the heap for the new data and it will not cause a runtime error. 
	// Bit weird !

	// create a channel of ints with buffer size 5 using make
	c := make(chan int, 5)
	c <- 1
	c <- 2
	fmt.Println("Channel:", c)
}