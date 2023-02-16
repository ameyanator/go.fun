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
 * In go we can mainly define errors on the fly with `errors.New()` but if you want extra information about the error you can define your own custom error
 * A smart idea could be to define a package with error variables of the common errors you exepct and then use it
 * Use panics when something unexpected happens, don't panic unless there's no clear way to handle the situation, use errors instead
 * Use `defer()` and `recover()` from panics that can occur when making further calls within a function
 * Always close files, network connections, sockets etc.

 ## Logging
 * Decide what flags to use when creating a logger
 * You can log to a network resource which will keep on writing on the socket. **We should use event streaming for logging**
 * To handle back pressure when writing to a network resource one might consider converting the protocol from tcp to udp. Pros and Cons are those of TCP vs UDP
 * We should decide which level to write to when using syslog
 * `runtime` and `runtime/debug` packages contain numerous functions for analyzing memory usage, goroutines, threading and other aspects of program's memory usages
 * When logging a stack we can set flags to print out stacks for all running goroutines. Tremendously useful when debugging concurrency problems, but increases output size.
 * We can print stacks from `runtime` & `runtime/debug`. When using `runtime` we must decide ahead how much buffer we need to allocate.
 * For advanced key-value pair logging which is easily understood by computers, [slog](https://pkg.go.dev/github.com/gookit/slog#section-readme) package exists. [Talk to understand slog](https://opensourcelive.withgoogle.com/events/go-day-2022/watch?talk=talk2)

 ## Http, Rest and Web Services
 * When connecting to an external server connection timeouts may occur due to a variety of reasons, it would make sense to retry the request if we are able to accurately detect a connection timeout
 * When downloading large files, connection timeouts may occur, in these cases don't start the download again from the start but rather continue downloading it from where you left off
 * If your service needs to have its own custom logic of errors, then return custom erros. Document it well so your clients know about it so they can parse it easily at their end.
 * Suppose some service has not documented the custom errors well, handle it by parsing unknown json structs
 * Have proper API versioning in URL or content type

 ## Cloud Computing
 * Working with multiple cloud providers. One must not get tied down to a single cloud provider, have interfaces in place and use it instead of concrete implementations in your code to avoid vendor lock-in
 * For error handling, have a look at common errors that occur and return them as such so that your native code handles the errors correctly irrespective of cloud provider error
 * Gather as much information possible about the host as applications on the cloud run in multiple different environments
 * As applications can run in different environments some of them might have dependencies preinstalled and others not, detect the dependencies before using them. if not then might even consider code to write the dependencies.
 * Cross Compile for multiple architectures and check your code's sanity on them
 * Monitor the runtime to get an idea into goroutine issues, memory leaks and other problems.