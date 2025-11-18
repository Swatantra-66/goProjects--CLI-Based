// Marshaling
package main

import (
	"encoding/json"
	"net/http"
)

type Dog struct {
	DogName      string          `json:"dogName"`
	DogAge       uint8           `json:"dogAge"`
	DogBreed     string          `json:"dogBreed"`
	DogBehaviour map[string]bool `json:"dogBehaviour"`
	DogFavFoods  []string        `json:"favFoods"`
}

func main() {
	dog := Dog{
		DogName:      "Don",
		DogAge:       10,
		DogBreed:     "German Shephard",
		DogBehaviour: map[string]bool{"Friendly": true, "Angry": false},
		DogFavFoods:  []string{"Milk", "Chicken", "Rice", "Biscuits", "Roti"},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		encoder, err := json.MarshalIndent(dog, "", " ")
		// err := encoder.Encode(dog)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(encoder)

	})
	http.ListenAndServe(":8080", mux)
}
