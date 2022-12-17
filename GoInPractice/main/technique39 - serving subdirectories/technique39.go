package main

import "net/http"

/*
A website or web application built with Go doesn’t need to sit behind a web server.
Instead, it handles serving all of the content with its web server, whether that content
is application pages or static files, such as CSS, images, or JavaScript.
*/

/*
Although Go is typically run as a server that serves all content, it can be used with a
Common Gateway Interface (CGI) or FastCGI server. The net/http/cgi package
works with the CGI interface, and the net/http/fastcgi package works with a
FastCGI interface. In this environment, static content may be served by another web
server. These packages are intended for compatibility with existing systems. CGI
starts up a new process to respond to each request, which is less efficient than typ-
ical Go serving. This is a setup we don’t recommend using.
*/

func main() {
	// dir := http.Dir("./files/")
	// handler := http.StripPrefix("/static/", http.FileServer(dir))
	// http.Handle("/static/", handler)
	// http.HandleFunc("/", homePage)
	// http.ListenAndServe(":8080", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
	homepage := "Home Page\n"
	res.Write([]byte(homepage))
	return
}
