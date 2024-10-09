package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/Jcastel2014/comments/internal/data"
)

func (a *applicationDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var importedData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&importedData)
	if err != nil {
		a.errorResponseJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", importedData)
}
