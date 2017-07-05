package main

import (
	"net"
	"bufio"
	"fmt"
	"io"
)

/*
Building upon the code from the previous problem:

Before we WRITE our RESPONSE , let's WRITE to our RESPONSE the STATUS LINE and some REPONSE HEADERS. Remember the request line and status line:

REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF https://tools.ietf.org/html/rfc7230#section-3.1.1

RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2

Write the following strings to the response - use io.WriteString for all of the following except the second and third:

"HTTP/1.1 200 OK\r\n"

fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))

fmt.Fprint(c, "Content-Type: text/plain\r\n")

"\r\n"

Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.
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
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if ln == "" {
			break
		}
	}
	body := "check the response via the network, hit f12"
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content-Length: %d\r\n",len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)
}
