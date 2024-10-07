package main

import "net/http"

func (a *applicationDependencies) createCommentHandler(w http.ResponseWriter, r *http.Request) {
	var incomingData struct {
		Content string `json:"content"`
		Author  string `json:"author"`
	}
}
