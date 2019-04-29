package main

import(
	"net"
	"log"
	"time"
	"bufio"
	"fmt"
)

func main() {
		go dial()
		listen()
}

func dial() {
	conn, err := net.Dial("tcp", ":9999")
	if err != nil {
		log.Print(err.Error())
	}

	text := []byte("here is some text")
	n, err := conn.Write(text)
	if err != nil {
		log.Print(err.Error())
	}
	log.Printf("Dialer: wrote %d bytes | text: %s", n, string(text))
}

func listen() {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Print(err.Error())
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err.Error())
		}

		go func(c net.Conn){
			bufReader := bufio.NewReader(c)
			defer c.Close()
			for {
				c.SetReadDeadline(time.Now().Add(time.Second * 10))

				bytes, err := bufReader.ReadBytes('\n')
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Printf("Listener read: %s", bytes)			
			}
		}(conn)	
	}
}