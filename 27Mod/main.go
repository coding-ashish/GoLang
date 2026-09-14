package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Mod lecture")

	//Declare a newRouter
	r := mux.NewRouter()

	//define a route and attach a handler function
	r.HandleFunc("/", serveHome).Methods("GET")

	//strt listening on port 4000
	//log.fatal automatically logs errors if the server fails to start
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	//w is used to send the response
	//r contaimns all the details of the incoming request (PARAMS, URLS , ETC..)
	w.Write([]byte("<h1>Hello Everyone ! My name Is Ashish Upadhyay</h1><br/><p>I am learning backend Development in GoLang.<p/>"))
}
