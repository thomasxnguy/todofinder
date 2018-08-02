TodoFinder 
===

todofinder is a tool that search for comment containing a specific key pattern such as "TODO" or "FIXME" in golang packages, and display results to the end-user with essential informations (files name, lines, comments) 

The tool has two execution mode :
- search mode : results are display in console stdout. It is the mode by default 
- server mode : service run as an API server on port 8080 by default


```
$ go run todofinder/cmd/todofinder.go -package fmt -pattern TODO 
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
```

Installation
---

The project uses "dep", a prototype dependency management tool for Go.
For more information, please visit https://github.com/golang/dep
```
% go get -u github.com/thomasxnguy/todofinder
% dep ensure (for Golang v1.9+)
```

Usage
---

```
$ todofinder -h
usage: todofinder <command>
The commands are:
   search - run in command line  (*default)
   server - run in server mode

search [-package input_package_name] [-pattern input_search_pattern]
  -package string
       package to search
  -pattern string
       patttern to search
```

### Command Line mode

```
$ go run todofinder/cmd/todofinder.go -h
```
or
```
$ go run todofinder/cmd/todofinder.go search -h
```

```
Usage of search:
  -package string
       package to search
  -pattern string
       patttern to search
```

Example
```
$ go run todofinder/cmd/todofinder.go -package fmt -pattern TODO 
/usr/local/go/src/fmt/scan.go : 740 :
TODO： accept N and Ni independently?
```

### Server mode

```
$ go run todofinder/cmd/todofinder.go server -h
```

```
Usage of server:
  -config string
    	configuration file path
```

#### Run in server mode
```
$ go run todofinder/cmd/todofinder.go server -config todofinder/conf/todofinder.yaml

```
The application runs as an HTTP server at port 8080 (default). It provides the following RESTful endpoints:

* `GET /search`: search a specific pattern in all .go file from a package

##### Example API request

```
$ go run todofinder/cmd/todofinder.go server -config todofinder/conf/todofinder.yaml &
$ curl -XGET 'localhost:8080/search?package=fmt&pattern=TODO'
{
    "result": [
        {
            "file": "/usr/local/Cellar/go/1.6.3/libexec/src/fmt/format.go",
            "pos": 332,
            "com": "TODO: Avoid buffer by pre-padding.\n"
        },
        {
            "file": "/usr/local/Cellar/go/1.6.3/libexec/src/fmt/scan.go",
            "pos": 747,
            "com": "TODO: accept N and Ni independently?\n"
        }
    ]
}

```

##### API Error codes


| Error Code | Description |
| --- | --- |
| NOT_FOUND | Resource was not found |
| METHOD_NOT_ALLOWED| Call endpoint using a not supported method |
| INTERNAL_SERVER_ERROR| An issue occurred server side |
| UNAUTHORIZED| User is not authorized to call this resource |
| BAD_PARAMETER| Call endpoint using a bad or missing parameter |
| PACKAGE_NOT_FOUND| Cannot find the package to search on |
| NO_SOURCE| The package does not contain valid source |
| SOURCE_NOT_READABLE| The package .go source files are not readable |


Packages
---

This service uses the following Go packages which can be easily replaced with your own favorite :

* FastHttp : Lightweight http server which is 10x faster than net/http. Benchmark and source are [here](https://github.com/valyala/fasthttp)
* Logrus : Structured logger for Go [here](https://github.com/sirupsen/logrus)
* Gingko/Gomega : BDD Testing framework for golang [here](https://onsi.github.io/ginkgo)
* Viper : Configuration solution for golang that can handle all types of formats and support watcher [here](https://github.com/spf13/viper)
* Ozzo-Validation : Configurable and extensible data validator [here](https://github.com/go-ozzo/ozzo-validation)
* Dep : Go dependency management tool [here](https://github.com/golang/dep)


Integration tests and benchmarks
---

This project contains a folder 'itest' containing integration tests and benchmark for todofinder service.
It needs 
* Baloo, an expressive and versatile end-to-end HTTP API testing made in Go : https://github.com/h2non/baloo
* Golang v1.7+

### Benchmark result

```
Benchmark_WrongEndpoint-4    	  500000	      3467 ns/op
Benchmark_BadParameters-4    	  500000	      3632 ns/op
Benchmark_BadMethod-4        	 1000000	      2332 ns/op
Benchmark_SearchEndpoint-4   	  500000	      2569 ns/op
```
