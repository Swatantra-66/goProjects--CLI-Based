package main

import (
	"bufio"
	"fmt"
	"os"
)

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What's your name?")
	if scanner.Scan() { //reads the next line.
		fmt.Printf("Hello, %v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input")
	}
}
