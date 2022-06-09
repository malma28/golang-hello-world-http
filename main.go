package main

import (
	"fmt"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello User!")
		return
	}
	fmt.Fprintf(w, "Hello %v!", name)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}

	http.HandleFunc("/", index)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: http.DefaultServeMux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
