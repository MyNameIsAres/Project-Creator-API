package main

import (
	"fmt"
	"net/http"

	// ? This may be completely wrong, and an unecessary, use of the . operator.
	. "github.com/org/project_creator_api/middleware"
)

// ? Is it get? Is it post? Is it .. gost? Pet? Who knows.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Whoooo lives in a pinaple under the sea")
}

func runServer() {
	mux := http.NewServeMux()
	myHandler := http.HandlerFunc(handler)

	//! Test routes, do not include in production.
	mux.Handle("/api/project", Middleware(myHandler, "log", "filter", "foobar"))
	mux.Handle("/api/project/foobar", MiddlewareGroup(myHandler, "api"))

	http.ListenAndServe(":8080", mux)
}

func main() {
	runServer()
}
