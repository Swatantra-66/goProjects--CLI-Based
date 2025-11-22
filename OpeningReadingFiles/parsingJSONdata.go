package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/sanity-io/litter"
)

type userData struct {
	ID        string         `json:"id"`
	UserName  string         `json:"username"`
	Email     string         `json:"email"`
	IsActive  bool           `json:"isActive"`
	Roles     []string       `json:"roles"`
	Profile   map[string]any `json:"profile"`
	LastLogin string         `json:"lastLogin"`
}

func parseJSON() {
	file, err := os.ReadFile("userData.json")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	var user userData
	err = json.Unmarshal(file, &user)
	if err != nil {
		log.Fatalf("Unable to parse JSON due to %s.", err)
	}

	litter.Dump(user)
}
