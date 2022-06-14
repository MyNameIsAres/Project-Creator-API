package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/org/project_creator_api/models"
	"github.com/org/project_creator_api/services/validate"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
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

	err = validate.Validate(project)
	if err != nil {
		fmt.Println(err)
	}
}
