package main

import (
	"log"
	"os"
)

func appendFiles() {
	// Open file with append mode flag
	file, err := os.OpenFile("goDocs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Unable to append to the file due to %s.", err)
	}

	defer file.Close()

	// Write data to the file
	data := "Hello, this is appended text!\n"
	bytesWritten, err := file.WriteString(data)
	if err != nil {
		log.Fatalf("Unable to write to the file due to %s.", err)
	}

	log.Printf("Successfully appended %d bytes to goDocs.txt\n", bytesWritten)
}
