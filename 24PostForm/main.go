package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Welcome to the Postform code :)")
	PerformPostformRequest()
}

func PerformPostformRequest() {
	const myUrl = "https://httpbin.org/post"

	data := url.Values{}
	data.Add("Name", "Ashish")
	data.Add("Garvit", "Padai krenge bai")
	data.Add("Rounak", "Bakk Saale")
	data.Add("Lord", "Mai Hurmesha Ati ati krta rehta hu")
	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	toPrint, error := io.ReadAll(response.Body)

	if error != nil {
		panic(error)
	}

	fmt.Println("The response we get after the Postform request is : ", string(toPrint))

}
