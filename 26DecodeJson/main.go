package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Decoding of JSon")
	decodeJson()
}

func decodeJson() {
	// Mock data mimicking a response from the web
	jsonDataFromWeb := []byte(`
        {
            "coursename": "ReactJS Bootcamp",
            "Price": 299,
            "website": "LearnCodeOnline.in",
            "tags": ["web dev", "js"]
        }
    `)

	validity := json.Valid(jsonDataFromWeb)

	if validity {
		var lcocourse course //Declare a variable of type course struct

		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json data is invalid")
	}
}
