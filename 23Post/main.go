package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to the POST lectore of GoLang")
	PerformPostRequest()
}

func PerformPostRequest() {
	const myUrl = "https://httpbin.org/post" // When I included "/post" in the URL then only the Post Request worked. Wiothout this it was throwing error.
	requestBody := strings.NewReader(`
    {
        "course" : "Golang",
		"price" : 0,
		"Desc" : "Hello Everyone! Welcome to the course"
    }
    `)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	toPrint, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("The response of the POST request is : ", string(toPrint))
}