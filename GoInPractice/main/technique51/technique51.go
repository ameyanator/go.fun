package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Custom HTTP error parsing

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Message  string `json:"message"`
}

func JSONError(w http.ResponseWriter, e Error) {
	data := struct {
		Err Error `json:"error"`
	}{e}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	fmt.Fprint(w, string(b))
}

func displayError(w http.ResponseWriter, r *http.Request) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     123,
		Message:  "An error occured",
	}
	JSONError(w, e)
}

func main() {
	http.HandleFunc("/", displayError)
	http.ListenAndServe(":8080", nil)
}
