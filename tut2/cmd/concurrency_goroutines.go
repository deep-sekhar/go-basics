// learn concurrency in go
package main

import (
	"fmt"
	"time"
	"sync"
)

// we can achieve concurrency in go using goroutines and channels
var wg = sync.WaitGroup{}
// wait groups are used to wait for a collection of goroutines to finish executing
var m = sync.Mutex{}
// mutexes are used to protect shared resources from concurrent access
// go also provides read-write mutexes (sync.RWMutex) which allow multiple readers but only one writer at a time
var rwm = sync.RWMutex{}

var dbData = []string{"data1", "data2", "data3", "data4", "data5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go dbCall(i) // start a new goroutine for each DB call
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Printf("Time taken: %v\n", time.Since(t0))
}

// simulate a DB call
func dbCall(i int) {
	fmt.Println("DB call %d started", i)
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("DB call %d finished", i)

	// using a regular mutex - only thread holding the lock can access the critical section of code, other threads will be blocked until the lock is released
	// m.Lock()
	// results = append(results, dbData[i]) // save the result to the shared resource (results slice)
	// m.Unlock()

	save(dbData[i]) // save the result to the shared resource (results slice)
	log() // log the current results -> read lock doesn't block other goroutines from writing to results
	wg.Done() 
	// signal that this goroutine has finished
}

func save(result string) {
	rwm.Lock()
	results = append(results, result)
	fmt.Printf("Added string: %s\n", result)
	rwm.Unlock()
}

func log() {
	rwm.RLock()
	// read lock for reading results
	fmt.Printf("Current results: %v\n", results)
	rwm.RUnlock()
}