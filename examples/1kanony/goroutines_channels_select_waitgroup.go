package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	selectExample()
	waitGroupExample()
}

func selectExample() {
	channel1 := make(chan interface{})
	channel2 := make(chan interface{})

	go func() {
		channel1 <- 1
		close(channel1)
	}()

	go func() {
		channel2 <- "random"
		close(channel2)
	}()

	closed1, closed2 := false, false

	for {
		select {
		case <-channel1: fmt.Println("Received msg from channel1"); closed1 = true
		case <-channel2: fmt.Println("Received msg from channel2"); closed2 = true
		default:
			fmt.Println("No traffic")
		}

		if closed1 && closed2 {
			break
		}
	}

	time.Sleep(time.Second)
}

func waitGroupExample() {
	names := []string{ "Akif", "Rza", "Memmed" }
	var group sync.WaitGroup

	group.Add(len(names))

	for _, name := range names {
		go doSomeStuff( name, &group )
	}

	group.Wait()
}

func doSomeStuff (name string, group *sync.WaitGroup) {
	defer group.Done()

	fmt.Printf( "%s has finished their stuff\n", name )
}
