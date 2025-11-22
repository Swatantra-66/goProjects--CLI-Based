package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func copyData() {
	file, err := os.Open("userData.json")
	if err != nil {
		log.Fatalf("Unable to open file userData.json due to %s", err)
	}

	defer file.Close()

	newFile, err := os.Create("copyUserData.json")
	if err != nil {
		log.Fatalf("Unable to create file due to %s.", err)
	}

	defer newFile.Close()

	n, err := io.Copy(newFile, file)
	if err != nil {
		log.Fatalf("Unable to copy file due to %s.", err)
	}

	fmt.Printf("Copied %d bytes.\n", n)
}
