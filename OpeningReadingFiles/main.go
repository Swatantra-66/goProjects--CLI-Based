// opening and reading files
package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {

	// Launch multiple goroutines to read concurrently
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			file, err := os.OpenFile("goDocs.txt", os.O_RDONLY, 0644)
			if err != nil {
				log.Fatalf("Unable to open file due to %s.\n", err)
			}
			defer file.Close() //task: Ensure file closed at end

			data := make([]byte, 500)
			count, err := file.Read(data)
			if err != nil {
				log.Fatalf("Unable to read file due to %s.\n", err)
			}

			fmt.Printf("Goroutine %d read %d bytes: %s\n", id, count, string(data[:count]))
		}(i)
	}

	wg.Wait() // Wait for all goroutines to finish

	writeFiles()       //task: writing files
	appendFiles()      //task: appending files
	parseJSON()        //task: parsing json data
	workingDirectory() //task: working with directory
	copyData()         //task: copying data
	userInput()        //task: reading user input
	buffer()           //task: collects data in a buffer and writes it in chunks, improving performance.
}
