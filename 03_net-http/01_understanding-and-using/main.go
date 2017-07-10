package main

/*====================================================================================
type Handler
====================================================================================
A Handler responds to an HTTP request.

ServeHTTP should write reply headers and data to the ResponseWriter and then return.
Returning signals that the request is finished; it is not valid to use the
ResponseWriter or read from the Request.Body after or concurrently
with the completion of the ServeHTTP call.

Depending on the HTTP client software, HTTP protocol version, and any intermediaries
between the client and the Go server, it may not be possible to read from the
Request.Body after writing to the ResponseWriter. Cautious handlers should
read the Request.Body first, and then reply.

Except for reading the body, handlers should not modify the provided Request.

If ServeHTTP panics, the server (the caller of ServeHTTP) assumes that the
effect of the panic was isolated to the active request. It recovers the
panic, logs a stack trace to the server error log, and hangs up the
connection. To abort a handler so the client sees an interrupted
response but the server doesn't log an error, panic with the
value ErrAbortHandler.

type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}

Any type that has ServeHTTP method is implementing the type handler

*/

import (
	"fmt"
	"net/http"
)

type anything int // create a type

func (input anything) ServeHTTP(w http.ResponseWriter, r *http.Request) { //that type is implementing type handler
	fmt.Fprintln(w, "Any code you want in this func")
}

func main() {
	var d anything

	/*=======================================================
	func ListenAndServe
	=======================================================
	func ListenAndServe(addr string, handler Handler) error

	ListenAndServe listens on the TCP network address addr
	and then calls Serve with handler to handle
	requests on incoming connections.
	Accepted connections are
	configured to enable
	TCP keep-alives.

	Handler is typically nil, in which case the
	DefaultServeMux is used.
	*/

	http.ListenAndServe(":8080", d) // so the handler can be passed onto the http package
}

//run this, open localhost:8080 from your browser
