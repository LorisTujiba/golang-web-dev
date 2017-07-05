package main

import (
	"net"
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
Building upon the code from the previous problem:

Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in

tags.
 */

func main(){

	lis,err := net.Listen("tcp",":8080")
	if err != nil{
		panic(err)
	}

	defer lis.Close()

	for{
		conn,err := lis.Accept()
		if err != nil{
			panic(err)
		}
		go serve(conn)
	}

}

func serve (conn net.Conn){
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	var i int
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0{
			xs := strings.Fields(ln)
			method := xs[0]
			uri := xs[1]
			fmt.Println("Method \t\t\t: ",method)
			fmt.Println("URI \t\t\t: ",uri)
		}

		if ln == "" {
			break
		}
		i++
	}
	body := `
	<h1>Holy cow this is low level</h1>`
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content-Length: %d\r\n",len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)
}
