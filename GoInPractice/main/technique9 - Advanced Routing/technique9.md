## PROBLEM
The built-in http package isn’t flexible enough, or doesn’t perform well in a particular use case.

## Solution
Routing URLs to functions is a common problem for web applications. Therefore, numerous packages have been built and tested, and are commonly used to tackle the problem of routing. A common technique is to import an existing request router and use it within your application.

Popular solutions include the following:
* github.com/julienschmidt/httprouter is considered a fast routing package with a focus on using a minimal amount of memory and taking as little time as possible to handle routing. It has features such as the ability to have case insensitive paths, cleaning up /../ in a path, and dealing with an optional trailing /.

* github.com/gorilla/mux is part of the Gorilla web toolkit. This loose collection of packages provides components you can use in an application. The mux package provides a versatile set of criteria to perform matching against, including host, schemes, HTTP headers, and more.

* github.com/bmizerany/pat provides a router inspired by the routing in Sinatra. The registered paths are easy to read and can contain named parameters such as /user/:name. It has inspired other packages such as github.com/gorilla/pat.