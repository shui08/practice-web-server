package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "howdy")
}

func main() {
	// http.Dir() takes in the name of a directory as a string and then returns
	// a FileSystem object. http.FileServer then accepts this FileSystem object
	// and returns a Handler object. this Handler object is able to respond to
	// http requests. so essentially, as the variable name suggests, we have
	// created an object that serves files from the "./static" directory
	// over HTTP.
	fileServer := http.FileServer(http.Dir("./static"))

	// this line sets fileServer to be the default handler for
	// `GET /(root path or any subpath)` requests.
	http.Handle("/", fileServer)

	// these two lines set the formHandler and helloHandler functions to handle
	// `GET /form` and `GET /hello` requests, respectively.
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// these lines are basically what starts the server. ListenAndServe starts
	// an HTTP server and tells it to listen on port 8080. the second argument
	// in ListenAndServe specifies what will be used to handle oncoming
	// requests. in this case, nil means that `DefaultServeMux` will be the
	// handler. this matches patterns to their corresponding function.
	// ListenAndServe also returns an error, which will be `nil` if the server
	// successfully starts. if there is an error, log.Fatal(err) will be called,
	// which prints an error message and terminates the program.
	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}