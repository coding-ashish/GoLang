package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://google.com:8080/html?name=Ashish&course=GoLang"

func main() {
	fmt.Println("Welcome to url handling.")
	result, err := url.Parse(myURL)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	Qparams := result.Query()
	fmt.Printf("The type of Qparams is : %T", Qparams)

	for _, val := range Qparams {
		fmt.Println(val)
	}

	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/imghp",
		RawQuery: "hl=en&ogbl",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)
}
