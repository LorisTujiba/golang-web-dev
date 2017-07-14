package _1

import (
	"net/http"
	"fmt"
	"context"
	"log"
	"time"
)

/*
==============================================================================
Context
==============================================================================
defines the context type which carries deadlines, cancelation signals,
and other request scoped values across API boundaries and between
processes.

Context makes it possible to manage a chain of calls within the same
call path by signaling context's Done channel.
*/

var err error

func main(){

	http.HandleFunc("/",foo)
	http.HandleFunc("/bar",bar)
	http.ListenAndServe(":8080",nil)

}

func bar(w http.ResponseWriter,r *http.Request){
	ctx := r.Context()

	log.Println(ctx)
	fmt.Fprintln(w,ctx)
}

func foo(w http.ResponseWriter,r *http.Request){
	ctx := r.Context()

	ctx = context.WithValue(ctx,"userID", 777)
	ctx = context.WithValue(ctx,"fName", "Bond")

	results := dbAccess(ctx)// pull out the data using this dbAccess to pass the data back
	results2,err := dbAccess2(ctx)
	if err!=nil{
		panic(err)
	}

	fmt.Fprintln(w,results)
	fmt.Fprintln(w,results2)

}

func dbAccess2(ctx context.Context)(int,error){
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	ch:=make(chan int)
	go func(){
		uid := ctx.Value("userID").(int)

		if ctx.Err() != nil{
			return
		}

		ch <- uid
	}()

	select {
	case <- ctx.Done():
		return 0, ctx.Err()
	case i := <-ch:
		return i, nil
	}

}

func dbAccess(ctx context.Context) int{
	uid := ctx.Value("userID").(int)
	return uid
}

//per request variables
//good candidate for putting into context, like the session id