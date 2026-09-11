package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the json lecture")
	JsonInGo()
}

type Course struct {
	Name     string   `json:"coursename"` //we can give alias to the struct
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`              // to make password invisible
	Tags     []string `json:"tags,omitempty"` //omitempty is used to specify the Golang that we don't need to omit tags if it is nil or null
}

func JsonInGo() {

	Courses := []Course{
		{"Javascript", 49, "Youtube", "abcd123", []string{"Good course", "A rated"}},
		{"GoLang", 59, "Instagram", "efgh123", []string{"Okay course", "B rated"}},
		{"FrontEnd", 19, "Facebook", "ijkl123", nil},
	}

	Json, err := json.MarshalIndent(Courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(Json))
}
