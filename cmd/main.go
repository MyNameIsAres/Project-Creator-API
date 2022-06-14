package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// ? This may be completely wrong, and an unecessary, use of the . operator.
	. "github.com/org/project_creator_api/internal/validate"
	"github.com/org/project_creator_api/models"
)

type Foo struct {
	Bar    string `json:"bar" validate:"required"`
	Tester string `json:"Tester" validate:"required"`
}

type BarEtje struct {
}

// ? Is it get? Is it post? Is it .. gost? Pet? Who knows.

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Request failed. Please create a POST request only.", http.StatusMethodNotAllowed)
	}
	requestBody, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var project models.Project
	err = json.Unmarshal(requestBody, &project)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = Validate(project)
	if err != nil {
		fmt.Println(err)
	}

	// template, err := services.CreateTemplate("main_class.tmpl", "./internal/templates/")
	// fmt.Println(template)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func runServer() {
	mux := http.NewServeMux()

	myHandler := http.HandlerFunc(postHandler)

	//! Test routes, do not include in production.
	mux.Handle("/foo", myHandler)
	// mux.Handle("/api/project", middleware.Middleware(myHandler, "log", "filter", "foobar"))
	http.ListenAndServe(":8080", mux)
}

func main() {
	runServer()
}
