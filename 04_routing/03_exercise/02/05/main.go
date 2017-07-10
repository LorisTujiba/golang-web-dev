package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

/*
Add code to WRITE to the connection.
*/

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go serve(conn)
	}

}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	io.WriteString(conn, "This is the response")
}
