package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start_time0 := time.Now().UnixMilli()
	wordstextbytes, err := os.ReadFile("unsorted10k.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	wordList := strings.Split(string(wordstextbytes), "\n")
	//lets sort
	start_time1 := time.Now().UnixMilli()
	sorted := mergeSort(wordList)
	sort_time := time.Now().UnixMilli() - start_time1
	fmt.Println("File loaded", start_time1-start_time0, "milli seconds")
	fmt.Println("Sorted in", sort_time, "milli seconds")

	for {
		var key string
		fmt.Println("Enter word or prefix to search:")
		fmt.Scanln(&key)
		s, e := findWordsPrefix(sorted, key, 5)
		if e <= s {
			fmt.Println("Netice tapilmadi")
			continue
		}
		fmt.Println("Netice:")
		for i := s; i < e; i++ {
			fmt.Println((sorted[i]))
		}
	}

}
