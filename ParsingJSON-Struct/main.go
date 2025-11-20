// Unmarshaling
package main

import (
	"encoding/json"
	"log"

	"github.com/sanity-io/litter"
)

type Dog struct {
	DogName      string          `json:"dogName"`
	DogAge       uint8           `json:"dogAge"`
	DogBreed     string          `json:"dogBreed"`
	DogBehaviour map[string]bool `json:"dogBehaviour"`
	FavFoods     []string        `json:"favFoods"`
}

func main() {
	input := `{
	"DogName" : "Don",
	"DogAge": 8,
	"DogBreed" : "German Shepherd",
	"DogBehaviour" : {
	"Friendly" : true, 
	"Fiery": false
	},
	"FavFoods" : [
	"Chicken",
	"Egg",
	"Milk",
	"Roti",
	"Biscuits",
	"Rice"
	 ]
	}`

	var dog Dog

	err := json.Unmarshal([]byte(input), &dog)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON due to %s.", err)
	}

	litter.Dump(dog)
}
