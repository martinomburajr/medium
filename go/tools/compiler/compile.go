package main

import (
	"io"
	"log"
	"os"
)

func main() {
	_, err := io.WriteString(os.Stdout, "Here is some output")
	if err != nil {
		log.Print("error writing %v", err)
	}
}
