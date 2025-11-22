package main

import (
	"fmt"
	"io"
	"os"
)

func workingDirectory() {
	err := os.Mkdir("newWorkingDirectory", 0755)
	if err != nil {
		fmt.Println("Error working with directory: ", err)
	}

	dirEntries, err := os.ReadDir(".")
	if err != nil && err != io.EOF {
		fmt.Println("Error reading directory: ", err)
	}

	for _, dirEntry := range dirEntries {
		fmt.Println(dirEntry.Name())
	}
}
