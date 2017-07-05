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

Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

Add this data to your RESPONSE so that this data is displayed in the browser.
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
	body := "check the response via the network, hit f12"
	io.WriteString(conn,"HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn,"content-Length: %d\r\n",len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn,"\r\n")
	io.WriteString(conn,body)
}
