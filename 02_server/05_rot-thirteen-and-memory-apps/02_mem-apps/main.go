package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"strings"
	"fmt"
)

func main(){
	lis,err := net.Listen("tcp",":8083")
	if err != nil{
		log.Fatal(err)
	}
	defer lis.Close()

	for{
		conn, err := lis.Accept()
		if err!= nil{
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn){
	defer conn.Close()

	io.WriteString(conn,"USE:\n" +
		"SET key value\n" +
		"GET key value\n" +
		"DEL key value\n")

	//read and write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan(){
		ln := scanner.Text()
		fs := strings.Fields(ln)//split by space

		switch fs[0] {// switch the first word
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintln(conn,value)
		case "SET":
			if len(fs) != 3{// check if the inputted command id three words
				fmt.Fprintln(conn,"Three words!")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value
			fmt.Fprintln(conn,"Added!\n")
		case "DEL":
			key := fs[1]
			delete(data,key)
			fmt.Fprintln(conn,"Deleted!\n")
		default:
			fmt.Fprintln(conn,"Invalid command")
		}
	}
}
