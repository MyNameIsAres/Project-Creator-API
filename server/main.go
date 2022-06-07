package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	// ? This may be completely wrong, and an unecessary, use of the . operator.
	"github.com/org/project_creator_api/middleware"
)

// ? Is it get? Is it post? Is it .. gost? Pet? Who knows.

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request failed. Please create a POST request only.", http.StatusMethodNotAllowed)
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	fmt.Println(string(requestBody))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}

func runServer() {
	mux := http.NewServeMux()

	myHandler := http.HandlerFunc(postHandler)

	//! Test routes, do not include in production.
	mux.Handle("/api/project", middleware.Middleware(myHandler, "log", "filter", "foobar"))
	http.ListenAndServe(":8080", mux)
}

func main() {
	runServer()
}
