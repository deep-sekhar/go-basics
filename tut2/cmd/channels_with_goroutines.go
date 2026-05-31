package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    toyChannel := make(chan string)
    bikeChannel := make(chan string)
    toys := []string{"Toy Car", "Doll"}

    var prodWg sync.WaitGroup
    prodWg.Add(len(toys) * 2) // two producers per toy

    for _, toy := range toys {
        go func(t string) { checkToyPrice(t, toyChannel); prodWg.Done() }(toy)
        go func(t string) { checkBikePrice(t, bikeChannel); prodWg.Done() }(toy)
    }

    // close channels when all producers are done
    go func() {
        prodWg.Wait()
        close(toyChannel)
        close(bikeChannel)
    }()

    // receive until both channels are closed
    for toyChannel != nil || bikeChannel != nil {
        select {
        case t, ok := <-toyChannel:
			// ok flag indicates if the channel is closed or not, if it is closed, we set the channel to nil so that it will not be selected in the select statement, and we continue to the next iteration of the loop
            if !ok {
                toyChannel = nil
                continue
            }
            fmt.Printf("Found a message on toy : %s\n", t)
        case b, ok := <-bikeChannel:
            if !ok {
                bikeChannel = nil
                continue
            }
            fmt.Printf("Found a message on bike : %s\n", b)
        }
    }
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