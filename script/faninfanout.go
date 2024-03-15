package script

import (
	"fmt"
	"sync"
	"time"
)

func generator(nums ...int) <-chan int {
	myChannel := make(chan int) //declare a channel

	go func() {
		//iterate the nums data and sends it to channel
		for _, val := range nums {
			myChannel <- val
		}
		close(myChannel)
	}()

	return myChannel
}

func FanInOut() {
	data1 := []int{1, 2, 3, 4, 5}
	data2 := []int{10, 20, 30, 40, 50}
	var wg sync.WaitGroup

	//it receives a "receive-only" directional channel
	ch1 := generator(data1...)
	ch2 := generator(data2...)
	wg.Add(2)

	//we will loop through both the channels till all data is sent and marked as close
	go func() {
		for val := range ch1 {
			fmt.Printf("Channel1 data: %v\n", val)
		}
		wg.Done()
	}()

	go func() {
		for val := range ch2 {
			fmt.Printf("Channel2 data: %v\n", val)
		}
		wg.Done()
	}()

	wg.Wait() //will wait till the above goroutines are marked as done
}

type Item struct {
	ID            int
	Name          string
	PackingEffort time.Duration
}

func PrepareItems(done <-chan bool) <-chan Item {
	items := make(chan Item)
	itemsToShip := []Item{
		Item{0, "Shirt", 1 * time.Second},
		Item{1, "Legos", 1 * time.Second},
		Item{2, "TV", 5 * time.Second},
		Item{3, "Bananas", 2 * time.Second},
		Item{4, "Hat", 1 * time.Second},
		Item{5, "Phone", 2 * time.Second},
		Item{6, "Plates", 3 * time.Second},
		Item{7, "Computer", 5 * time.Second},
		Item{8, "Pint Glass", 3 * time.Second},
		Item{9, "Watch", 2 * time.Second},
	}
	go func() {
		for _, item := range itemsToShip {
			select {
			case <-done:
				return
			case items <- item:
			}
		}
		close(items)
	}()
	return items
}

func PackItems(done <-chan bool, items <-chan Item, workerID int) <-chan int {
	packages := make(chan int)
	go func() {
		for item := range items {
			select {
			case <-done:
				return
			case packages <- item.ID:
				time.Sleep(item.PackingEffort)
				fmt.Printf("Worker #%d: Shipping package no. %d, took %ds to pack\n", workerID, item.ID, item.PackingEffort/time.Second)
			}
		}
		close(packages)
	}()
	return packages
}

func merge(done <-chan bool, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingPackages := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case outgoingPackages <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingPackages)
	}()
	return outgoingPackages
}

func FanInOut2() {
	done := make(chan bool)
	defer close(done)

	start := time.Now()

	items := PrepareItems(done)

	workers := make([]<-chan int, 4)
	for i := 0; i < 4; i++ {
		workers[i] = PackItems(done, items, i)
	}

	numPackages := 0
	for range merge(done, workers...) {
		numPackages++
	}

	fmt.Printf("Took %fs to ship %d packages\n", time.Since(start).Seconds(), numPackages)
}
