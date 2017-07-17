package main

import (
	"encoding/base64"
	"fmt"
)

/*======================================================
64 Base Encoding
======================================================
Usualy used if you have double quote on your value
that you want to store into a cookie or session.
so before you put it inside the cookie or sess,
you can use the 64 based encoding.
*/

func main() {

	s := "Love love is but a song to sing fear's the way we die. You can make the" +
		"mountains ring or make tha angels cry."

	encodedStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234efghijklmnopqrstuvwxyz!@#$%^&*()+/" // you can set your own encoding so that the result is only consist of this chars
	s64 := base64.NewEncoding(encodedStd).EncodeToString([]byte(s))

	//or you can use the standard encoding

	s64b := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)
	fmt.Println(s64b)

}
