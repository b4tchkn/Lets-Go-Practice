package main

import (
	"fmt"
	"net/http"
)

func handler(weiter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(weiter, "Hello, %s", request.URL.Path[1:])
}

func main() {

}