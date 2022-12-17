package main

import "net/http"

func main() {
	http.HandleFunc("/", readme)
	http.ListenAndServe(":8080", nil)
}

func readme(res http.ResponseWriter, req *http.Request) {
	/*
		ServeFile looks at the
		If-Modified-Since HTTP header and responds with a 304 Not Modified response
		if possible.
	*/
	http.ServeFile(res, req, "readme.txt")
}
