// writing files
package main

import (
	"log"
	"os"
)

func writeFiles() {
	file, err := os.OpenFile("goDocs.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Unable to open file due to %s.", err)
	}

	defer file.Close()

	data := []byte("Learning Golang\n")
	bytesWritten, err := file.Write(data)
	if err != nil {
		log.Fatalf("Unable to write file due to %s.", err)
	}

	log.Printf("Wrote %d bytes to goDocs.txt\n", bytesWritten)
}
