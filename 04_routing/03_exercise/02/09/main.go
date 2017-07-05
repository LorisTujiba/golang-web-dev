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

Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply
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
	var method,uri string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0{
			xs := strings.Fields(ln)
			method = xs[0]
			uri = xs[1]
			fmt.Println("Method \t\t\t: ",method)
			fmt.Println("URI \t\t\t: ",uri)
		}

		if ln == "" {
			break
		}
		i++
	}



	switch {
	case method == "GET" && uri == "/":
		io.WriteString(conn,"This is index")
	case method == "GET" && uri == "/apply":
		io.WriteString(conn,"get from apply")
	case method == "POST" && uri == "/apply":
		io.WriteString(conn,"This is posting")
	default:
		io.WriteString(conn,"this is default")
	}
}

func preHandle(c net.Conn){
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
}