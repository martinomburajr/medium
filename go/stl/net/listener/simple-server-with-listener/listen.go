package main

/*
*	This file uses net.Listen to listen to connections on the TCP network on port 9999. It uses a bufio to read the connection, and uses concurrency for each new connection the listener listens to. If sending an HTTP request the current code will only read up to the empty line, and pause for the duration of the readTimeOut set in the acceptConnections method.
*/

import(
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
	"io"
)

func main() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Print(err.Error())
	}

	defer func() {
		listener.Close()
		fmt.Println("Listener closed")
	}()

	for {
		conn, err := listener.Accept()
		log.Print("Accepting Connection on TCP Port 9999\n")
		if err != nil {
			log.Fatalf(err.Error())
		}

		go acceptedConnection(conn)
	}
}

// acceptedConnection takes in a connection and reads the connection using a buffered reader. Each invokation of this method will start its own infinite loop that only exits when the buffered reader fails to read more bytes
func acceptedConnection(conn net.Conn) {
	bufReader := bufio.NewReader(conn)
	defer conn.Close()

	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		data, err := bufReader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		fmt.Printf("Listener Read: %s", data)
	}
	log.Print("goroutine exited connection")
}
