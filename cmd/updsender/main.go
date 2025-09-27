package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Error: ", err)
	}

	conn, err := net.DialUDP("localhost:42069", nil, addr)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error: ", err)
		}

		if _, err := conn.Write([]byte(line)); err != nil {
			log.Println("Error:", err) // FIX: use Println or a %v verb
		}
	}

}
