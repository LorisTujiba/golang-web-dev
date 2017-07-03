package main

/*======================================================
Rotate 13
======================================================
is a simple letter substitution cipher that replaces a
letter with the letter 13 letters after it in the
alphabet. ROT13 is a special case of the Caesar
cipher, developed in ancient Rome.
*/

import (
	"net"
	"log"
	"bufio"
	"strings"
	"fmt"
)

func main(){

	lis,err := net.Listen("tcp",":8080")
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
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := strings.ToLower(scanner.Text())
		bs := []byte(ln)
		r := rot13(bs)

		fmt.Fprintf(conn,"%s - %s",ln,r)
	}
}

func rot13(input []byte) []byte{ // create another slice of byte
	var r13 = make([]byte, len(input))
	for i, v := range input{//get index and value
		//ascii 97-122 (a-z)
		if v <= 109{
			r13[i] = v + 13 //if lower than m, add 13 char
		}else{
			r13[i] = v - 13	// if lower, subtract 13 char
		}
		return r13
	}

	return input
}
