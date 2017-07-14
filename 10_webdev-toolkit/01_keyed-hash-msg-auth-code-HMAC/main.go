package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
)

/*======================================================
HMAC
======================================================
How hash work? The same input produced the same output.
This hash will need a key, so it will produce
different output if the key is different.
even though the thing you hashd is the
same.
 */

func main(){

	hash("test@example.com")
	hash("test@example.com")
	hash("test@examplw.com")

}

func hash(s string){
	h := hmac.New(sha256.New,[]byte("myKey"))
	io.WriteString(h,s)
	fmt.Println(h.Sum(nil))
}