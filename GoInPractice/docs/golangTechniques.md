 # Golang Techniques Guide

 ## General
 * One can easily work with configuration files in json, yaml and ini.  - technique 3
 * OS package helps to get values from environment - technique 4
 
 ## HTTP and Web Servers
 * We should use channels and goroutines to make graceful shutdowns when dealing with web/http servers
 * HTTP Routing is a big topic and there are many ways to deal with it, some famous libraries/frameworks exist (httprouter, gorrila framework [technique 9]) and some simple default ways are there (technique 6-8)

 ## Concurrency and GoRoutines

 ## Error Handling
 * We should minimize returning nils from programs and return default values whenever possible
 * In go we can mainly define errors on the fly with errors.New() but if you want extra information about the error you can define your own custom error
 * A smart idea could be to define a package with error variables of the common errors you exepct and then use it
 * Use panics when something unexpected happens, don't panic unless there's no clear way to handle the situation, use errors instead
 * Use `defer()` and `recover()` from panics that can occur when making further calls within a function
 * Always close files, network connections, sockets etc.