package api

import (
	"fmt"
	"net/http"
)

func (a *APIHandler) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
	}
}
