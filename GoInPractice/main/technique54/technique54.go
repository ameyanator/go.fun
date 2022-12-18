package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API version in URL

/*
APIs are typically versioned by major number changes such as v1, v2, and v3. This
number scheme signifies breaking changes to the API. An application designed to work
with v2 of an API won’t be able to consume the v3 API version because it’s too different.
*/

/*
But what about API changes that add functionality to an existing API? For example,
say that functionality is added to the v1 API. In this case, the API can be incremented
with a point version; feature additions can increment the API to v1.1. This tells devel-
opers and applications about the additions.
*/

/*
Provide the API version in the REST API URL. For example, instead of providing an API
of https://example.com/api/todos, add a version to the path so it looks like
https://example.com/api/v1/todos.
*/

type testMessage struct {
	Message string `json:"message"`
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test Message"}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", displayTest)
	http.ListenAndServe(":8080", nil)
}
