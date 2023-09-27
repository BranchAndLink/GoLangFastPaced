package main

import (
	"fmt"
	"time"
)

//With Goroutines, concurrency is achieved in Go programming. It helps two or more independent functions to run together.

// create a function
func display(message string) {

	fmt.Println(message)
}

func MyRoutine(söz string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(i, ")", söz)
		//fasilə
		time.Sleep(2 * time.Second)
	}

}

func main() {

	// call goroutine
	// go display("Process 1")

	display("starting")
	time.Sleep(time.Second * 3)
	go display("Process 1")
	go display("Process 2")
	go display("Process 3")

	// to sleep main goroutine for 1 sec
	// lst := []string{""}

	var sözlər = [...]string{"Salam", "Baku", "Mars", "Yupiter"}
	for _, s := range sözlər {
		go MyRoutine(s, 2)
	}
	time.Sleep(3 * time.Second)

	// create channel of integer type
	number := make(chan int)
	//number <- 15
	// access type and value of channel
	fmt.Printf("Channel Type: %T\n", number)
	fmt.Printf("Channel Value: %v\n", number)

	message := make(chan string)

	// function call with goroutine
	go channelData(number, message)

	// retrieve channel data
	fmt.Println("Channel string data:", <-number)
	fmt.Println("Channel  Data:", <-message)

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)
	//we are calling sendData() before receiving data from the channel. That's why the first "No receiver..." is printed
	// receive channel data
	fmt.Println(<-ch)

	// function call with goroutine
	go receiveData(ch)

	fmt.Println("No data. Receive Operation Blocked")

	// send data to the channel
	ch <- "Data Received. Receive Operation Successful"

}

func channelData(number chan int, message chan string) {

	// send data into channel
	number <- 15
	message <- "Learning Go channel"

}

func sendData(ch chan string) {

	// data sent to the channel
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}

func receiveData(ch chan string) {

	// receive data from the channel
	fmt.Println(<-ch)

}
