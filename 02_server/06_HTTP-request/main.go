package main

import (
	"net"
	"bufio"
	"fmt"
	"strings"
)

func main(){

	lis,err := net.Listen("tcp",":8088")
	if err!= nil{
		panic(err)
	}

	defer lis.Close()

	for{
		conn,err := lis.Accept()
		if err!=nil{
			panic(err)
		}

		go handle(conn)
	}

}

func handle(conn net.Conn){
	defer conn.Close()

	// read request
	request(conn)

	//write response
	respond(conn)
}

func request(conn net.Conn){
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0{
			//request line
			m := strings.Fields(ln)[0]//method
			u := strings.Fields(ln)[1]//uri
			fmt.Println("Method ",m)
			fmt.Println("Url ",u)

			//using multiplexer
			mux(conn,ln)
		}
		if ln == ""{
			//headers are done
			break
		}
		i++
	}
}

func respond(conn net.Conn){

	fmt.Println(conn,"responded")

}

func mux(conn net.Conn, ln string){
	m := strings.Fields(ln)[0]//method
	u := strings.Fields(ln)[1]//uri
	fmt.Fprintln(conn,"Method ",m)
	fmt.Fprintln(conn,"Url ",u)

	//multiplexer
	if m == "GET" && u == "/"{
		fmt.Fprintln(conn,"\n Index \n")
	}
	if m == "GET" && u == "/about"{
		fmt.Fprintln(conn,"\n About \n")
	}
	if m == "GET" && u == "/profile"{
		fmt.Fprintln(conn,"\n Profile \n")
	}
	if m == "POST" && u == "/apply"{
		fmt.Fprintln(conn,"\n Applying \n")
	}

}