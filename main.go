package main

import (
	"fmt"
	"log"
	"net/http"
)

// this function is a handler for the "/hello" endpoint of the server. it takes
// in a response writer, which indicates where responses will be written, and
// a pointer to a request. if the url path that the request was made on is not
// "/hello", then the function displays a 404 not found error and terminates.
// additionally, if r is not a "GET" request, then it once again displays an
// error and terminates. if both of these conditions are passed, then the
// function simply displays hello to the ResponseWriter w.
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

// this function is a handler for the "/form" endpoint of the server. when a
// user submits their name and address into the "/form.html" endpoint, it then
// transfers them to the "/form" endpoint and submits a POST request with their
// name and address. we then parse the request, which fills out the "r.Form" and
// "r.PostForm" fields. ParseForm also returns an error, so we check if the
// error is non nil. if it is, then we print the error message and return.
// if not, we retrieve the user input from "name" and "address" and display it.
func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parsing error: %v", err)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintln(w, "POST request successful!")
	fmt.Fprintf(w, "Name: %v\n", name)
	fmt.Fprintf(w, "Address: %v", address)
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
