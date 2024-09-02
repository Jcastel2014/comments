package main

import (
	"fmt"
	"net/http"
)

func (a *applicationDependencies) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status:available")
	fmt.Fprintln(w, "environment: %s\n", a.config.environment)
	fmt.Fprintln(w, "version: %s\n", appVersion)
}
