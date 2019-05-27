package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// t - is a time variable we show on entry.
var t time.Time
func init() {
	t = time.Now()
}

// main is the application entrypoint
func main() {
	// We Create a server listening on PORT 8080, and we pass in our Router
	log.Printf("Server started on PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", Router()))
}

// Router is a simmple http.ServeMux multiplexer that returns a Handler (note http.DefaultServeMux is a Handler)
func Router() *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", BaseHandler)

	return mux
}

// BaseHandler handles any request to the / endpoint.
// In our example it checks to see the HTTP Method is a GET and if not it returns a HTTP 400 error, otherwise,
// it prints a simple message.
func BaseHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("ensure you use HTTP Method GET"))
	}

	writer.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Hey I am doing some work under the hood. I was initialized at %s", t.String())
	writer.Write([]byte(msg))
}