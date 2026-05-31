// channels are a powerful feature in Go that allow goroutines to communicate and synchronize with each other
// 1. channels hold data of a specific type and can be used to send and receive values between goroutines

// 2. channels can be buffered or unbuffered
// unbuffered channels block the sender until the receiver is ready to receive the value, and block the receiver until a value is sent
// buffered channels allow the sender to send values without blocking until the buffer is full, and allow the receiver to receive values without blocking until the buffer is empty

// 3. channels can be directional (send-only or receive-only) or bidirectional (can send and receive)

// 4. channels can be closed to indicate that no more values will be sent on the channel, and receivers can check if a channel is closed using the comma-ok idiom

// 5. channels can be used with select statements to wait on multiple channel operations simultaneously, allowing for more complex synchronization patterns

package main

import (
	"fmt"
	"time"
	"math/rand"
)

// use channels with goroutines to achieve concurrency and synchronization
func main() {
	// var c = make(chan int) // this is a bidirectional and unbuffered channel of type int
	// go process(c)
	// for i:= range c { // receive values from the channel until it is closed
	// 	fmt.Printf("Received value: %d\n", i)
	// }

	// buffer channel
	var c2 = make(chan int, 5) // this is a bidirectional and buffered channel of type int with buffer size 5
	go process2(c2)
	for i := range c2 {
		fmt.Printf("Received value: %d\n", i)
	}
	// here we are using a buffered channel, so the sender can send values without blocking until the buffer is full, and the receiver can receive values without blocking until the buffer is empty. In prev one we were using an unbuffered channel, so the sender and receiver were blocking each other until they were ready to send and receive values, which is not the case here.

	var toyChannel = make(chan string) 
	var bikeChannel = make(chan string)
	var toys = []string{"Toy Car", "Doll"}
	for _, toy := range toys {
		go checkToyPrice(toy, toyChannel) // start a new goroutine for each toy to check its price
		go checkBikePrice(toy, bikeChannel) // start a new goroutine for each bike to check its price
	}
	// sendMessage(toyChannel) // start a new goroutine to receive messages from the toyChannel
	sendBothMessages(toyChannel, bikeChannel) // start a new goroutine to receive messages from both channels
}

func process(c chan int) {
	fmt.Println("Processing started")
	time.Sleep(2 * time.Second)
	fmt.Println("Processing finished")
	for i:=0 ; i<5; i++ {
		c <- i // send values to the channel
	}
	close(c) // close the channel to signal that no more values will be sent
	// if we iterate like for i:= range c it will keep waiting for values to be sent on the channel until it is closed
	// so if we dont close the channel, the for loop will keep waiting for values to be sent on the channel and it will never exit, which will cause a deadlock
}

// demo for buffered channels
func process2(c chan<- int) {
	defer close(c) // close the channel when the function returns 
	for i:=0 ; i<5; i++ {
		c <- i // send values to the channel
	}
	fmt.Println("Processing finished")
}

var maxPrice = 50.0 // default type is float64
func checkToyPrice(toy string, toyChannel chan string) {
	for {
		time.Sleep(1 * time.Second) // simulate checking the price every 1 second
		var price = rand.Float64() * 100
		if price < maxPrice {
			toyChannel <- toy // send the toy name to the channel if the price is less than the max price
			break
		}
	}
}

func checkBikePrice(bike string, bikeChannel chan string) {
	for {
		time.Sleep(1 * time.Second) // simulate checking the price every 1 second
		var price = rand.Float64() * 100
		if price < maxPrice {
			bikeChannel <- bike // send the bike name to the channel if the price is less than the max price
			break
		}
	}
}

// func sendMessage(toyChannel chan string) {
// 	fmt.Println("Found a message on toy : %s\n", <-toyChannel) // receive a message from the channel and print it
// }
// you can also write this function as:
// func sendMessage(toyChannel chan string) {
// 	for toy := range toyChannel { // receive messages from the channel until it is closed
// 		fmt.Printf("Found a message on toy : %s\n", toy)
// 	}
// }
// here we are using a for loop to receive messages from the channel until it is closed, which is more efficient than receiving a single message and then exiting the function, because it allows us to receive multiple messages from the channel without having to start a new goroutine for each message. In the previous version of the function, we were only receiving a single message from the channel and then exiting the function, which is not ideal if we want to receive multiple messages from the channel.

func sendBothMessages(toyChannel, bikeChannel chan string) {
	select {
		case toy := <-toyChannel: // receive a message from the toyChannel
			fmt.Printf("Found a message on toy : %s\n", toy)
		case bike := <-bikeChannel: // receive a message from the bikeChannel
			fmt.Printf("Found a message on bike : %s\n", bike)
	}
}

// Note: Internally for non-buffered channels, when sender is ready to send but there are no receivers then sender is moved from ready queue to waiting queue and it is blocked until a receiver is ready to receive the value, and when a receiver is ready to receive but there are no senders then receiver is moved from ready queue to waiting queue and it is blocked until a sender is ready to send a value. For buffered channels, when sender is ready to send but the buffer is full then sender is moved from ready queue to waiting queue and it is blocked until there is space in the buffer, and when a receiver is ready to receive but the buffer is empty then receiver is moved from ready queue to waiting queue and it is blocked until there is a value in the buffer.