package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	dial, err := net.Dial("tcp", ":8080") // dial to a server with 8080 as the port
	if err != nil {
		log.Fatal(err)
	}
	defer dial.Close()

	fmt.Fprintln(dial, "Aku udah masuk nih") //write to sever

}
