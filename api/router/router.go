package router

import (
	"net/http"

	handler "github.com/org/project_creator_api/api/handler"
)

func ServeRouter() {
	mux := http.NewServeMux()

	myHandler := http.HandlerFunc(handler.PostHandler)

	mux.Handle("/foo", myHandler)
	// mux.Handle("/api/project", middleware.Middleware(myHandler, "log", "filter", "foobar"))
	http.ListenAndServe(":8080", mux)
}
