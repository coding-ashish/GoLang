package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {

	fmt.Println("Welcome to web requests")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	fmt.Printf("The response is of type %T\n:", response)

	defer response.Body.Close() //It is user's responsibility to close the request.

	databytes, error := io.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}
	fmt.Println(string(databytes))
	
}