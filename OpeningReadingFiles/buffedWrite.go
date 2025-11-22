package main

import (
	"bufio"
	"log"
	"os"
)

func buffer() {
	file, err := os.Create("user.txt")
	if err != nil {
		log.Fatalf("Unable to create user.txt due to %s.", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("User is Swatantra")
	writer.Flush()
}
