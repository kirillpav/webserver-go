package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(ch)

		str := ""

		for {
			data := make([]byte, 8)
			n, err := f.Read(data)

			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				ch <- str
				str = ""
			}

			str += string(data)
		}
		if len(str) != 0 {
			ch <- str
			// continue
		}
	}()

	return ch
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	defer listener.Close()

	// infinite loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error occured: %s", err)
		}

		for line := range getLinesChannel(conn) {
			fmt.Printf("read: %s", line)
		}
	}
}
