package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:    ":7000",
		Handler: http.DefaultServeMux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
