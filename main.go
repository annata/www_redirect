package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", redirect)
	server := &http.Server{Addr: ":80", Handler: nil}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func redirect(response http.ResponseWriter, request *http.Request) {
	host := request.Host
	if host == "" || strings.HasPrefix(host, "www") {
		response.WriteHeader(404)
		response.Write(([]byte)("not fount"))
	} else {
		response.Header().Set("Location", "http://www."+host)
		response.WriteHeader(301)
	}
}
