package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { //return a bool every time it scan. When no longer given us something it'll be false
		//automatically break by line
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

//can also run it using web browser, type the localhost:8080
//and the request will be printed here
//or type yourself using another cmd as the client
