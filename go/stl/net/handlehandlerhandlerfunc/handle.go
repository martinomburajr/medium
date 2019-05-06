package handlehandlerhandlerfunc

import (
"fmt"
"log"
"net/http"
)

// switch top-level package to package main in order to run
func main() {
	x := SomeTypeThatImplHandler{
		Text: "Here is some text!!",
	}
	http.Handle("/", x)

	log.Fatal(http.ListenAndServe(":8080", x))
}

// SomeTypeThatImplHandler can contain whatever you want. And conents can be accessed in ServeHTTP
type SomeTypeThatImplHandler struct {
	Text string
}

// ServeHTTP is a hard requirement for a Handler. Which is a requirement as the 2nd argument for the Handle.
func (x SomeTypeThatImplHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	str := fmt.Sprintf("%s | %s", x.Text,  "Note we didnt need a Server Mux")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(str)); err != nil {
		return
	}
}