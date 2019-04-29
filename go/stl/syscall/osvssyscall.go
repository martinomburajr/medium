package medium_syscall

import (
	"log"
	"os"
	"syscall"
)

// OSOpenRead performs a read operation on a given file.
// The filepath is a generic filepath to the file.
// It takes in a buffer that can be preset and returns the number of bytes
// read.
func OSOpenRead(filepath string, buffer int) int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	data :=  make([]byte, buffer)
	n, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// SyscallOpenRead performs a read operation on a given file using the direct syscall.
// The filepath is a generic filepath to the file.
// It takes in a buffer that can be preset and returns the number of bytes
// read.
func SyscallOpenRead(filepath string, buffer int) int {
	fd, err := syscall.Open(filepath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	data :=  make([]byte, buffer)
	n, err := syscall.Read(fd, data)
	if err != nil {
		log.Fatal(err)
	}

	err = syscall.Close(fd)
	if err != nil {
		log.Fatal(err)
	}
	return n
}