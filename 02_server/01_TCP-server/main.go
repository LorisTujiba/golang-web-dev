package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8088") //Listen asks for two string
	//1.what kind of network you want to listen on ?
	//2.what port ?
	//Listener is an interface that gave back 3 methods, accept, close, and address

	if err != nil {
		log.Fatal(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() //if somebody comes in, we accept. Returns in connection
		//connection provide a read and write functions
		//so if you have a connection, you can choose to listen to it
		//or you can write to it, like "hello!"
		if err != nil {
			log.Fatal(err)
		}

		/*
			Since conn is a writer, implements the writer interface
			we can pass that connection in println, printf
			because they're once a writer
		*/
		io.WriteString(conn, "\nHello, im the server\n")
		fmt.Fprintln(conn, "\nWhat's up?")
		fmt.Fprintf(conn, "%v", "\nWhat's up yo?")

		conn.Close() //close the connection and wait for the next request
	}

	//to run the server, turn on the telnet
	//then open cmd, go run main.go
	//then open another cmd
	//telnet localhost 8088<- number of the port

}
